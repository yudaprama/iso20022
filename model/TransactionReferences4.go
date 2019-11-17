package model

// Identifies the underlying transaction.
type TransactionReferences4 struct {

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
	EndToEndIdentification *Max35Text `xml:"EndToEndId"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level.
	// Usage: The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId,omitempty"`

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`
}

func (t *TransactionReferences4) SetPaymentInformationIdentification(value string) {
	t.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences4) SetInstructionIdentification(value string) {
	t.InstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences4) SetEndToEndIdentification(value string) {
	t.EndToEndIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences4) SetTransactionIdentification(value string) {
	t.TransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences4) SetMandateIdentification(value string) {
	t.MandateIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences4) AddCreditorSchemeIdentification() *PartyIdentification43 {
	t.CreditorSchemeIdentification = new(PartyIdentification43)
	return t.CreditorSchemeIdentification
}
