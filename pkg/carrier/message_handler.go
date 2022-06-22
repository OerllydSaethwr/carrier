package carrier

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
)

func (c *Carrier) handleInitMessage(rawMessage message.Message) error {
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

	c.broadcast(echoM)

	c.locks.ValueStore.Lock()
	c.stores.valueStore[h] = initM.V
	c.locks.ValueStore.Unlock()

	return nil
}

func (c *Carrier) handleEchoMessage(rawMessage message.Message) error {
	var err error
	echoM, ok := rawMessage.(*message.EchoMessage)
	if !ok {
		return fmt.Errorf("expected EchoMessage")
	}

	err = c.Verify(echoM.H, echoM.S)
	if err == nil {
		c.locks.SignatureStore.Lock()
		c.stores.signatureStore[echoM.H] = append(c.stores.signatureStore[echoM.H], echoM.S)

		// TODO potential deadlock, but I think it's fine for now
		// @Critical C2
		if len(c.stores.signatureStore[echoM.H]) == c.f+1 {
			c.locks.SuperBlockSummary.Lock()
			c.stores.superBlockSummary.payload[echoM.H] = c.stores.signatureStore[echoM.H]
			c.locks.SuperBlockSummary.Unlock()
		}

		c.locks.SignatureStore.Unlock()
	}

	c.locks.SuperBlockSummary.Lock()
	if len(c.stores.superBlockSummary.payload) == c.n-c.f {
		err = c.NestedPropose(c.stores.superBlockSummary)
		c.refreshSuperBlock()
	}
	c.locks.SuperBlockSummary.Unlock()
	if err != nil {
		return err
	}

	return nil
}

func (c *Carrier) handleRequestMessage(rawMessage message.Message) error {
	requestM, ok := rawMessage.(*message.RequestMessage)
	if !ok {
		return fmt.Errorf("expected RequestMessage")
	}

	c.locks.ValueStore.Lock()
	defer c.locks.ValueStore.Unlock()
	if v, ok := c.stores.valueStore[requestM.H]; ok {
		resolveM := message.NewResolveMessage(
			requestM.H,
			v,
			c.GetID(),
		)
		dest := c.neighbours[requestM.GetSenderID()]
		dest.marshalAndSend(resolveM)
	}

	return nil
}

func (c *Carrier) handleResolveMessage(rawMessage message.Message) error {
	resolveM, ok := rawMessage.(*message.ResolveMessage)
	if !ok {
		return fmt.Errorf("expected ResolveMessage")
	}

	hb := sha256.Sum256(resolveM.Payload())
	h := hex.EncodeToString(hb[:])
	if resolveM.H == h {
		c.locks.ValueStore.Lock()
		c.stores.valueStore[h] = resolveM.V
		c.locks.ValueStore.Unlock()
	}

	c.locks.AcceptedHashStore.Lock()
	if _, ok := c.stores.acceptedHashStore.payload[h]; ok {
		c.stores.acceptedHashStore.payload[h] = resolveM.V

	}
	c.locks.AcceptedHashStore.Unlock()

	c.checkAcceptedHashStoreAndDecide()

	return nil
}

// handleNestedSMRDecision assumes that N is safe and correct
func (c *Carrier) handleNestedSMRDecision(N SuperBlockSummary) {
	// Adding a bottleneck here because of time constaints
	// We should be able to process Nested SMR decisions concurrently
	// This requires us to keep a separate instance of D for each decision we receive
	// I didn't have time to implement this so I'm ensuring that we process each D sequentially

	c.locks.DecisionLock.Lock()
outer:
	for h, S := range N.payload {
		if len(S) != c.f+1 {
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
		c.locks.AcceptedHashStore.Lock()
		c.locks.ValueStore.Lock()
		if _, ok := c.stores.valueStore[h]; ok {
			c.stores.acceptedHashStore.payload[h] = c.stores.valueStore[h]

			// Unlock
			c.locks.ValueStore.Unlock()
			c.locks.AcceptedHashStore.Unlock()

		} else {
			c.stores.acceptedHashStore.payload[h] = nil

			// Unlock
			c.locks.ValueStore.Unlock()
			c.locks.AcceptedHashStore.Unlock()

			c.broadcast(message.NewRequestMessage(h, c.GetID()))
		}
	}

	c.checkAcceptedHashStoreAndDecide()
}
