package model

// Details of breakdown of a quantity.
type QuantityBreakdown39 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification39 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *Balance11 `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price3 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`

	// Valuation amounts for the lot provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts6 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts6 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts6 `xml:"AltrnRptgCcyAmts,omitempty"`
}

func (q *QuantityBreakdown39) AddLotNumber() *GenericIdentification39 {
	q.LotNumber = new(GenericIdentification39)
	return q.LotNumber
}

func (q *QuantityBreakdown39) AddLotQuantity() *Balance11 {
	q.LotQuantity = new(Balance11)
	return q.LotQuantity
}

func (q *QuantityBreakdown39) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown39) AddLotPrice() *Price3 {
	q.LotPrice = new(Price3)
	return q.LotPrice
}

func (q *QuantityBreakdown39) AddTypeOfPrice() *TypeOfPrice32Choice {
	q.TypeOfPrice = new(TypeOfPrice32Choice)
	return q.TypeOfPrice
}

func (q *QuantityBreakdown39) AddAccountBaseCurrencyAmounts() *BalanceAmounts6 {
	q.AccountBaseCurrencyAmounts = new(BalanceAmounts6)
	return q.AccountBaseCurrencyAmounts
}

func (q *QuantityBreakdown39) AddInstrumentCurrencyAmounts() *BalanceAmounts6 {
	q.InstrumentCurrencyAmounts = new(BalanceAmounts6)
	return q.InstrumentCurrencyAmounts
}

func (q *QuantityBreakdown39) AddAlternateReportingCurrencyAmounts() *BalanceAmounts6 {
	q.AlternateReportingCurrencyAmounts = new(BalanceAmounts6)
	return q.AlternateReportingCurrencyAmounts
}
