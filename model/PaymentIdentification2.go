package model

// Set of elements used to provide further means of referencing a payment transaction.
type PaymentIdentification2 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	//
	// Usage: the  instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification assigned by the initiating party to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to carry on multiple references, the end-to-end identification must be carried on throughout the entire end-to-end chain.
	EndToEndIdentification *Max35Text `xml:"EndToEndId"`

	// Unique identification assigned by the first instructing agent to unambiguously identify the transaction and passed on, unchanged, throughout the entire interbank chain.
	//
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level.
	// Usage: The instructing agent has to make sure the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId"`
}

func (p *PaymentIdentification2) SetInstructionIdentification(value string) {
	p.InstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentIdentification2) SetEndToEndIdentification(value string) {
	p.EndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentIdentification2) SetTransactionIdentification(value string) {
	p.TransactionIdentification = (*Max35Text)(&value)
}
