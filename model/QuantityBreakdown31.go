package model

// Details of breakdown of a quantity.
type QuantityBreakdown31 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`
}

func (q *QuantityBreakdown31) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown31) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}
