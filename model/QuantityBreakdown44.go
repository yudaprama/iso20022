package model

// Details of breakdown of a quantity.
type QuantityBreakdown44 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification39 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity15Choice `xml:"LotQty,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price3 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown44) AddLotNumber() *GenericIdentification39 {
	q.LotNumber = new(GenericIdentification39)
	return q.LotNumber
}

func (q *QuantityBreakdown44) AddLotQuantity() *FinancialInstrumentQuantity15Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown44) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	q.SecuritiesSubBalanceType = new(GenericIdentification47)
	return q.SecuritiesSubBalanceType
}

func (q *QuantityBreakdown44) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown44) AddLotPrice() *Price3 {
	q.LotPrice = new(Price3)
	return q.LotPrice
}

func (q *QuantityBreakdown44) AddTypeOfPrice() *TypeOfPrice32Choice {
	q.TypeOfPrice = new(TypeOfPrice32Choice)
	return q.TypeOfPrice
}
