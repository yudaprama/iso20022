package model

// Details of breakdown of a quantity.
type QuantityBreakdown33 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification39 `xml:"LotNb"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity15Choice `xml:"LotQty,omitempty"`
}

func (q *QuantityBreakdown33) AddLotNumber() *GenericIdentification39 {
	q.LotNumber = new(GenericIdentification39)
	return q.LotNumber
}

func (q *QuantityBreakdown33) AddLotQuantity() *FinancialInstrumentQuantity15Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.LotQuantity
}
