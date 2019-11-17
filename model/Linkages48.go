package model

// Information related to a linked transaction.
type Linkages48 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition10Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber6Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References58Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity4Choice `xml:"LkdQty,omitempty"`
}

func (l *Linkages48) AddProcessingPosition() *ProcessingPosition10Choice {
	l.ProcessingPosition = new(ProcessingPosition10Choice)
	return l.ProcessingPosition
}

func (l *Linkages48) AddMessageNumber() *DocumentNumber6Choice {
	l.MessageNumber = new(DocumentNumber6Choice)
	return l.MessageNumber
}

func (l *Linkages48) AddReference() *References58Choice {
	l.Reference = new(References58Choice)
	return l.Reference
}

func (l *Linkages48) AddLinkedQuantity() *PairedOrTurnedQuantity4Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity4Choice)
	return l.LinkedQuantity
}
