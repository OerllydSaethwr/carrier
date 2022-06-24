package carrier

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/superblock"
	"github.com/OerllydSaethwr/carrier/pkg/util"
)

func (c *Carrier) HandleInitMessage(rawMessage message.Message) error {
	initM, ok := rawMessage.(*message.InitMessage)
	if !ok {
		return fmt.Errorf("expected InitMessage")
	}

	h := initM.Hash()

	s := util.Signature{
		S:        c.Sign(h),
		SenderID: c.GetID(),
	}

	echoM := message.NewEchoMessage(
		h,
		s,
	)

	c.Broadcast(echoM)

	c.Locks.ValueStore.Lock()
	c.Stores.valueStore[h] = initM.V
	c.Locks.ValueStore.Unlock()

	return nil
}

func (c *Carrier) HandleEchoMessage(rawMessage message.Message) error {
	var err error
	echoM, ok := rawMessage.(*message.EchoMessage)
	if !ok {
		return fmt.Errorf("expected EchoMessage")
	}

	err = c.Verify(echoM.H, echoM.S)
	if err == nil {
		c.Locks.SignatureStore.Lock()
		c.Stores.signatureStore[echoM.H] = append(c.Stores.signatureStore[echoM.H], echoM.S)

		// TODO potential deadlock, but I think it's fine for now
		// @Critical C2
		if len(c.Stores.signatureStore[echoM.H]) == c.F+1 {
			c.Locks.SuperBlockSummary.Lock()
			c.Stores.superBlockSummary[echoM.H] = c.Stores.signatureStore[echoM.H]
			c.Locks.SuperBlockSummary.Unlock()
		}

		c.Locks.SignatureStore.Unlock()
	}

	c.Locks.SuperBlockSummary.Lock()
	if len(c.Stores.superBlockSummary) == c.N-c.F {
		err = c.NestedPropose(c.Stores.superBlockSummary)
		c.Stores.superBlockSummary = map[string][]util.Signature{}
		c.SbsCounter++ // Advance superblockcounter
	}
	c.Locks.SuperBlockSummary.Unlock()
	if err != nil {
		return err
	}

	return nil
}

func (c *Carrier) HandleRequestMessage(rawMessage message.Message) error {
	requestM, ok := rawMessage.(*message.RequestMessage)
	if !ok {
		return fmt.Errorf("expected RequestMessage")
	}

	c.Locks.ValueStore.Lock()
	defer c.Locks.ValueStore.Unlock()
	if v, ok := c.Stores.valueStore[requestM.H]; ok {
		resolveM := message.NewResolveMessage(
			requestM.H,
			v,
			c.GetID(),
		)
		dest := c.Neighbours[requestM.GetSenderID()]
		dest.MarshalAndSend(resolveM)
	}

	return nil
}

func (c *Carrier) HandleResolveMessage(rawMessage message.Message) error {
	resolveM, ok := rawMessage.(*message.ResolveMessage)
	if !ok {
		return fmt.Errorf("expected ResolveMessage")
	}

	hb := sha256.Sum256(resolveM.Payload())
	h := hex.EncodeToString(hb[:])
	if resolveM.H == h {
		c.Locks.ValueStore.Lock()
		c.Stores.valueStore[h] = resolveM.V
		c.Locks.ValueStore.Unlock()
	}

	c.Locks.AcceptedHashStore.Lock()
	if _, ok := c.Stores.acceptedHashStore[h]; ok {
		c.Stores.acceptedHashStore[h] = resolveM.V

	}
	c.Locks.AcceptedHashStore.Unlock()

	c.CheckAcceptedHashStoreAndDecide()

	return nil
}

// HandleNestedSMRDecision assumes that N is safe and correct
func (c *Carrier) HandleNestedSMRDecision(N superblock.SuperBlockSummary) {
	// Adding a bottleneck here because of time constaints
	// We should be able to process Nested SMR decisions concurrently
	// This requires us to keep a separate instance of D for each decision we receive
	// I didn't have time to implement this so I'm ensuring that we process each D sequentially

	c.Locks.DecisionLock.Lock()
outer:
	for h, S := range N {
		if len(S) != c.F+1 {
			continue outer
		}
		for _, s := range S {
			err := c.Verify(h, s)
			if err != nil {
				continue outer
			}
		}

		// This is ugly, but we need the locking to prevent concurrency bugs
		// Lock
		c.Locks.AcceptedHashStore.Lock()
		c.Locks.ValueStore.Lock()
		if _, ok := c.Stores.valueStore[h]; ok {
			c.Stores.acceptedHashStore[h] = c.Stores.valueStore[h]

			// Unlock
			c.Locks.ValueStore.Unlock()
			c.Locks.AcceptedHashStore.Unlock()

		} else {
			c.Stores.acceptedHashStore[h] = nil

			// Unlock
			c.Locks.ValueStore.Unlock()
			c.Locks.AcceptedHashStore.Unlock()

			c.Broadcast(message.NewRequestMessage(h, c.GetID()))
		}
	}

	c.CheckAcceptedHashStoreAndDecide()
}
