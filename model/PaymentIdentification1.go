package model

// Set of elements used to provide further means of referencing a payment transaction.
type PaymentIdentification1 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	//
	// Usage: the  instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification assigned by the initiating party to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	EndToEndIdentification *Max35Text `xml:"EndToEndId"`
}

func (p *PaymentIdentification1) SetInstructionIdentification(value string) {
	p.InstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentIdentification1) SetEndToEndIdentification(value string) {
	p.EndToEndIdentification = (*Max35Text)(&value)
}
