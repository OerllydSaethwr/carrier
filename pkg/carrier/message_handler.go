package carrier

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/xid"
)

func (c *Carrier) handleInitMessage(rawMessage message.Message) error {
	initM, ok := rawMessage.(*message.InitMessage)
	if !ok {
		return fmt.Errorf("expected InitMessage")
	}
	//log.Info().Msgf("received InitMessage from %s", initM.GetSenderID())

	h := initM.Hash()

	s := util.Signature{
		S:        c.Sign(h),
		SenderID: c.GetID(),
	}

	echoM := message.NewEchoMessage(
		h,
		s,
		c.GetID(),
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

		// TODO potential deadlock
		if len(c.stores.signatureStore[echoM.H]) == c.f+1 {
			newSBSum := SuperBlockSummaryItem{
				ID: xid.New().String(),
				H:  echoM.H,
				S:  c.stores.signatureStore[echoM.H], //TODO concurrent access here
			}
			c.locks.SuperBlockSummary.Lock()
			c.stores.superBlockSummary = append(c.stores.superBlockSummary, newSBSum)
			c.locks.SuperBlockSummary.Unlock()
		}

		c.locks.SignatureStore.Unlock()
	}

	c.locks.SuperBlockSummary.Lock()
	if len(c.stores.superBlockSummary) == c.n-c.f {
		err = c.NestedPropose(c.stores.superBlockSummary)
		c.stores.superBlockSummary = make([]SuperBlockSummaryItem, 0)
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
	if _, ok := c.stores.acceptedHashStore[h]; ok {
		c.stores.acceptedHashStore[h] = resolveM.V

	}
	c.locks.AcceptedHashStore.Unlock()

	return nil
}

func (c *Carrier) handleNestedSMRDecision(N SuperBlockSummary) error {
outer:
	for _, hs := range N {
		if len(hs.S) != c.f+1 {
			continue outer
		}
		for _, s := range hs.S {
			err := c.Verify(hs.H, s)
			if err != nil {
				continue outer
			}
		}

		// This is ugly, but we need the locking to prevent concurrency bugs
		// Lock
		c.locks.AcceptedHashStore.Lock()
		c.locks.ValueStore.Lock()
		if _, ok := c.stores.valueStore[hs.H]; ok {
			c.stores.acceptedHashStore[hs.H] = c.stores.valueStore[hs.H]

			// Unlock
			c.locks.ValueStore.Unlock()
			c.locks.AcceptedHashStore.Unlock()

		} else {
			c.stores.acceptedHashStore[hs.H] = nil

			// Unlock
			c.locks.ValueStore.Unlock()
			c.locks.AcceptedHashStore.Unlock()

			c.broadcast(message.NewRequestMessage(hs.H, c.GetID()))
		}
	}

	return nil
}
