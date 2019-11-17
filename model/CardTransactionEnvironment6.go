package model

// Environment of the transaction.
type CardTransactionEnvironment6 struct {

	// Institution initiator of the reconciliation (correspond to the ISO 8583 field 94).
	SendingInstitution *GenericIdentification73 `xml:"SndgInstn"`

	// Institution destination of the reconciliation (correspond to the ISO 8583 field 93).
	ReceivingInstitution *GenericIdentification73 `xml:"RcvgInstn"`
}

func (c *CardTransactionEnvironment6) AddSendingInstitution() *GenericIdentification73 {
	c.SendingInstitution = new(GenericIdentification73)
	return c.SendingInstitution
}

func (c *CardTransactionEnvironment6) AddReceivingInstitution() *GenericIdentification73 {
	c.ReceivingInstitution = new(GenericIdentification73)
	return c.ReceivingInstitution
}
