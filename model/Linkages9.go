package model

// Information related to a linked transaction.
type Linkages9 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition1Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References14Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity1Choice `xml:"LkdQty,omitempty"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification36Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages9) AddProcessingPosition() *ProcessingPosition1Choice {
	l.ProcessingPosition = new(ProcessingPosition1Choice)
	return l.ProcessingPosition
}

func (l *Linkages9) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages9) AddReference() *References14Choice {
	l.Reference = new(References14Choice)
	return l.Reference
}

func (l *Linkages9) AddLinkedQuantity() *PairedOrTurnedQuantity1Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity1Choice)
	return l.LinkedQuantity
}

func (l *Linkages9) AddReferenceOwner() *PartyIdentification36Choice {
	l.ReferenceOwner = new(PartyIdentification36Choice)
	return l.ReferenceOwner
}
