package model

// Information related to a linked transaction.
type Linkages38 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition7Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber5Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References47Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity3Choice `xml:"LkdQty,omitempty"`
}

func (l *Linkages38) AddProcessingPosition() *ProcessingPosition7Choice {
	l.ProcessingPosition = new(ProcessingPosition7Choice)
	return l.ProcessingPosition
}

func (l *Linkages38) AddMessageNumber() *DocumentNumber5Choice {
	l.MessageNumber = new(DocumentNumber5Choice)
	return l.MessageNumber
}

func (l *Linkages38) AddReference() *References47Choice {
	l.Reference = new(References47Choice)
	return l.Reference
}

func (l *Linkages38) AddLinkedQuantity() *PairedOrTurnedQuantity3Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity3Choice)
	return l.LinkedQuantity
}
