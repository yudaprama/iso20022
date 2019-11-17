package model

// Identification references of the underlying transaction.
type CertificateIdentification1 struct {

	// Point to point reference, as assigned by the instructing party of the underlying message.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Proprietary reference related to the underlying transaction.
	Proprietary *ProprietaryReference1 `xml:"Prtry,omitempty"`
}

func (c *CertificateIdentification1) SetMessageIdentification(value string) {
	c.MessageIdentification = (*Max35Text)(&value)
}

func (c *CertificateIdentification1) SetAccountServicerReference(value string) {
	c.AccountServicerReference = (*Max35Text)(&value)
}

func (c *CertificateIdentification1) SetPaymentInformationIdentification(value string) {
	c.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (c *CertificateIdentification1) SetInstructionIdentification(value string) {
	c.InstructionIdentification = (*Max35Text)(&value)
}

func (c *CertificateIdentification1) SetEndToEndIdentification(value string) {
	c.EndToEndIdentification = (*Max35Text)(&value)
}

func (c *CertificateIdentification1) AddProprietary() *ProprietaryReference1 {
	c.Proprietary = new(ProprietaryReference1)
	return c.Proprietary
}
