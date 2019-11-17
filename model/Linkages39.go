package model

// Information related to a linked transaction.
type Linkages39 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition8Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber5Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References46Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification92Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages39) AddProcessingPosition() *ProcessingPosition8Choice {
	l.ProcessingPosition = new(ProcessingPosition8Choice)
	return l.ProcessingPosition
}

func (l *Linkages39) AddMessageNumber() *DocumentNumber5Choice {
	l.MessageNumber = new(DocumentNumber5Choice)
	return l.MessageNumber
}

func (l *Linkages39) AddReference() *References46Choice {
	l.Reference = new(References46Choice)
	return l.Reference
}

func (l *Linkages39) AddReferenceOwner() *PartyIdentification92Choice {
	l.ReferenceOwner = new(PartyIdentification92Choice)
	return l.ReferenceOwner
}
