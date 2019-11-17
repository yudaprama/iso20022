package model

// Information related to a linked transaction.
type Linkages27 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition5Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References24Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification36Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages27) AddProcessingPosition() *ProcessingPosition5Choice {
	l.ProcessingPosition = new(ProcessingPosition5Choice)
	return l.ProcessingPosition
}

func (l *Linkages27) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages27) AddReference() *References24Choice {
	l.Reference = new(References24Choice)
	return l.Reference
}

func (l *Linkages27) AddReferenceOwner() *PartyIdentification36Choice {
	l.ReferenceOwner = new(PartyIdentification36Choice)
	return l.ReferenceOwner
}
