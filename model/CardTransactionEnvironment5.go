package model

// Environment of the transaction.
type CardTransactionEnvironment5 struct {

	// Institution initiator of the reconciliation.
	// It corresponds to the ISO 8583 field number 94.
	SendingInstitution *GenericIdentification73 `xml:"SndgInstn"`

	// Institution destination of the reconciliation.
	// It corresponds to the ISO 8583 field number 93.
	ReceivingInstitution *GenericIdentification73 `xml:"RcvgInstn"`

	// Institution in charge of the settlement of the transaction.
	SettlementInstitution *GenericIdentification73 `xml:"SttlmInstn"`
}

func (c *CardTransactionEnvironment5) AddSendingInstitution() *GenericIdentification73 {
	c.SendingInstitution = new(GenericIdentification73)
	return c.SendingInstitution
}

func (c *CardTransactionEnvironment5) AddReceivingInstitution() *GenericIdentification73 {
	c.ReceivingInstitution = new(GenericIdentification73)
	return c.ReceivingInstitution
}

func (c *CardTransactionEnvironment5) AddSettlementInstitution() *GenericIdentification73 {
	c.SettlementInstitution = new(GenericIdentification73)
	return c.SettlementInstitution
}
