package model

// Information related to a linked transaction.
type Linkages3 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition2Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References6Choice `xml:"Ref"`
}

func (l *Linkages3) AddProcessingPosition() *ProcessingPosition2Choice {
	l.ProcessingPosition = new(ProcessingPosition2Choice)
	return l.ProcessingPosition
}

func (l *Linkages3) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages3) AddReference() *References6Choice {
	l.Reference = new(References6Choice)
	return l.Reference
}
