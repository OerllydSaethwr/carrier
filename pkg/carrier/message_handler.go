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
	//log.Info().Msgf("received InitMessage from %s", initM.GetSender())

	h := initM.Hash()

	s := util.Signature{
		S:  c.Sign(h),
		Pk: c.GetStringPK(),
	}

	echoM := message.NewEchoMessage(
		h,
		s,
		c.GetAddress(),
	)

	c.broadcast(echoM)

	c.locks.ValueStore.Lock()
	c.valueStore[h] = initM.V
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
		c.signatureStore[echoM.H] = append(c.signatureStore[echoM.H], echoM.S)
		c.locks.SignatureStore.Unlock()

		if len(c.signatureStore[echoM.H]) == c.f+1 {
			newSBSum := SuperBlockSummaryItem{
				ID: xid.New().String(),
				H:  echoM.H,
				S:  c.signatureStore[echoM.H],
			}
			c.locks.SuperBlockSummary.Lock()
			c.superBlockSummary = append(c.superBlockSummary, newSBSum)
			c.locks.SuperBlockSummary.Unlock()
		}
	}

	c.locks.SuperBlockSummary.Lock()
	if len(c.superBlockSummary) == c.n-c.f {
		err = c.NestedPropose(c.superBlockSummary)
		c.superBlockSummary = make([]SuperBlockSummaryItem, 0)
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
	if v, ok := c.valueStore[requestM.H]; ok {
		resolveM := message.NewResolveMessage(
			requestM.H,
			v,
		)
		c.marshalAndSend(requestM.GetSender(), resolveM)
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
		c.valueStore[h] = resolveM.V
		c.locks.ValueStore.Unlock()
	}

	//TODO add locks
	if _, ok := c.acceptedHashStore[h]; ok {
		c.acceptedHashStore[h] = resolveM.V
	}

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

		if _, ok := c.valueStore[hs.H]; ok {
			c.acceptedHashStore[hs.H] = c.valueStore[hs.H]
		} else {
			c.acceptedHashStore[hs.H] = nil
			c.broadcast(message.NewRequestMessage(hs.H, c.GetAddress()))
		}
	}

	return nil
}
