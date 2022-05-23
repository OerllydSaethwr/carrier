package carrier

import (
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

func (c *Carrier) handleInitMessage(rawMessage message.Message) error {
	initM, ok := rawMessage.(*message.InitMessage)
	if !ok {
		return fmt.Errorf("expected InitMessage")
	}
	log.Info().Msgf("Received InitMessage from %s", initM.GetSender())

	h := initM.Hash()

	s := util.Signature{
		S:  c.Sign(h),
		Pk: c.GetStringPK(),
	}

	echoM := &message.EchoMessage{
		H:      h,
		S:      s,
		Sender: c.GetAddress(),
	}

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
	log.Info().Msgf("Received EchoMessage from %s", echoM.GetSender())

	err = c.Verify(echoM.H, echoM.S)
	if err == nil {
		c.locks.SignatureStore.Lock()
		c.signatureStore[echoM.H] = append(c.signatureStore[echoM.H], echoM.S)
		c.locks.SignatureStore.Unlock()

		if len(c.signatureStore[echoM.H]) == c.f+1 {
			newSBSum := SuperBlockSummaryItem{
				ID:         xid.New().String(),
				H:          echoM.H,
				Signatures: c.signatureStore[echoM.H],
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
	_, ok := rawMessage.(*message.RequestMessage)
	if !ok {
		return fmt.Errorf("expected RequestMessage")
	}

	log.Info().Msgf("Received RequestMessage")

	return nil
}

func (c *Carrier) handleResolveMessage(rawMessage message.Message) error {
	_, ok := rawMessage.(*message.ResolveMessage)
	if !ok {
		return fmt.Errorf("expected ResolveMessage")
	}

	log.Info().Msgf("Received ResolveMessage")

	return nil
}

func (c *Carrier) handleNestedSMRDecision(N SuperBlockSummary) {
	//for _, superBlockSummaryItem := range N {
	//	for _, s := range superBlockSummaryItem.Signatures {
	//		//bdn.Verify(c.suite, pk, superBlockSummaryItem.H, s)
	//	}
	//}
}
