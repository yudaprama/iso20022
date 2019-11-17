package model

// Details of breakdown of a quantity.
type QuantityBreakdown5 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *Number2Choice `xml:"LotNb"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`
}

func (q *QuantityBreakdown5) AddLotNumber() *Number2Choice {
	q.LotNumber = new(Number2Choice)
	return q.LotNumber
}

func (q *QuantityBreakdown5) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}
