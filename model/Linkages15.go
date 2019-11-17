package model

// Information related to a linked transaction.
type Linkages15 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber4Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *IdentificationReference8Choice `xml:"Ref"`
}

func (l *Linkages15) AddMessageNumber() *DocumentNumber4Choice {
	l.MessageNumber = new(DocumentNumber4Choice)
	return l.MessageNumber
}

func (l *Linkages15) AddReference() *IdentificationReference8Choice {
	l.Reference = new(IdentificationReference8Choice)
	return l.Reference
}
