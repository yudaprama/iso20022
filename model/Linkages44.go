package model

// Information related to a linked transaction.
type Linkages44 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition18Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber16Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References54Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification103Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages44) AddProcessingPosition() *ProcessingPosition18Choice {
	l.ProcessingPosition = new(ProcessingPosition18Choice)
	return l.ProcessingPosition
}

func (l *Linkages44) AddMessageNumber() *DocumentNumber16Choice {
	l.MessageNumber = new(DocumentNumber16Choice)
	return l.MessageNumber
}

func (l *Linkages44) AddReference() *References54Choice {
	l.Reference = new(References54Choice)
	return l.Reference
}

func (l *Linkages44) AddReferenceOwner() *PartyIdentification103Choice {
	l.ReferenceOwner = new(PartyIdentification103Choice)
	return l.ReferenceOwner
}
