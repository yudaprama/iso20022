package model

// Information related to a linked transaction.
type Linkages18 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber4Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *IdentificationReference11Choice `xml:"Ref"`
}

func (l *Linkages18) AddMessageNumber() *DocumentNumber4Choice {
	l.MessageNumber = new(DocumentNumber4Choice)
	return l.MessageNumber
}

func (l *Linkages18) AddReference() *IdentificationReference11Choice {
	l.Reference = new(IdentificationReference11Choice)
	return l.Reference
}
