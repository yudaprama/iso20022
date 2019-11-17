package model

// Choice between a turnaround and pair-off quantity for instructing a one to many and many to many (partial) pair-off or turnaround.
type PairedOrTurnedQuantity4Choice struct {

	// Quantity of financial instruments of the linked transaction to be paired-off.
	PairedOffQuantity *FinancialInstrumentQuantity15Choice `xml:"PairdOffQty,omitempty"`

	// Quantity of financial instruments of the linked transaction to be turned.
	TurnedQuantity *FinancialInstrumentQuantity15Choice `xml:"TrndQty,omitempty"`
}

func (p *PairedOrTurnedQuantity4Choice) AddPairedOffQuantity() *FinancialInstrumentQuantity15Choice {
	p.PairedOffQuantity = new(FinancialInstrumentQuantity15Choice)
	return p.PairedOffQuantity
}

func (p *PairedOrTurnedQuantity4Choice) AddTurnedQuantity() *FinancialInstrumentQuantity15Choice {
	p.TurnedQuantity = new(FinancialInstrumentQuantity15Choice)
	return p.TurnedQuantity
}
