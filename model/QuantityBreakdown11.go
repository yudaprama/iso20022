package model

// Details of breakdown of a quantity.
type QuantityBreakdown11 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTime1Choice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price4 `xml:"LotPric,omitempty"`
}

func (q *QuantityBreakdown11) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown11) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown11) AddLotDateTime() *DateAndDateTime1Choice {
	q.LotDateTime = new(DateAndDateTime1Choice)
	return q.LotDateTime
}

func (q *QuantityBreakdown11) AddLotPrice() *Price4 {
	q.LotPrice = new(Price4)
	return q.LotPrice
}
