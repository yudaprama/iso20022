package model

// Details of breakdown of a quantity.
type QuantityBreakdown3 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *Number2Choice `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown3) AddLotNumber() *Number2Choice {
	q.LotNumber = new(Number2Choice)
	return q.LotNumber
}

func (q *QuantityBreakdown3) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown3) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown3) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown3) AddTypeOfPrice() *TypeOfPrice3Choice {
	q.TypeOfPrice = new(TypeOfPrice3Choice)
	return q.TypeOfPrice
}
