package model

// Information related to a linked transaction.
type Linkages17 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition1Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References25Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity1Choice `xml:"LkdQty,omitempty"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification36Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages17) AddProcessingPosition() *ProcessingPosition1Choice {
	l.ProcessingPosition = new(ProcessingPosition1Choice)
	return l.ProcessingPosition
}

func (l *Linkages17) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages17) AddReference() *References25Choice {
	l.Reference = new(References25Choice)
	return l.Reference
}

func (l *Linkages17) AddLinkedQuantity() *PairedOrTurnedQuantity1Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity1Choice)
	return l.LinkedQuantity
}

func (l *Linkages17) AddReferenceOwner() *PartyIdentification36Choice {
	l.ReferenceOwner = new(PartyIdentification36Choice)
	return l.ReferenceOwner
}
