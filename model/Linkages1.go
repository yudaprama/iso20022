package model

// Information related to a linked transaction.
type Linkages1 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition1Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References1Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity1Choice `xml:"LkdQty,omitempty"`
}

func (l *Linkages1) AddProcessingPosition() *ProcessingPosition1Choice {
	l.ProcessingPosition = new(ProcessingPosition1Choice)
	return l.ProcessingPosition
}

func (l *Linkages1) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages1) AddReference() *References1Choice {
	l.Reference = new(References1Choice)
	return l.Reference
}

func (l *Linkages1) AddLinkedQuantity() *PairedOrTurnedQuantity1Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity1Choice)
	return l.LinkedQuantity
}
