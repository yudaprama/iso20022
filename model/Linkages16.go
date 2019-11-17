package model

// Information related to a linked transaction.
type Linkages16 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition2Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References24Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification36Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages16) AddProcessingPosition() *ProcessingPosition2Choice {
	l.ProcessingPosition = new(ProcessingPosition2Choice)
	return l.ProcessingPosition
}

func (l *Linkages16) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages16) AddReference() *References24Choice {
	l.Reference = new(References24Choice)
	return l.Reference
}

func (l *Linkages16) AddReferenceOwner() *PartyIdentification36Choice {
	l.ReferenceOwner = new(PartyIdentification36Choice)
	return l.ReferenceOwner
}
