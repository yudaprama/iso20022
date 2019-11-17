package model

// Details of breakdown of a quantity.
type QuantityBreakdown28 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *Balance7 `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`

	// Valuation amounts for the lot provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts2 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts2 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts2 `xml:"AltrnRptgCcyAmts,omitempty"`
}

func (q *QuantityBreakdown28) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown28) AddLotQuantity() *Balance7 {
	q.LotQuantity = new(Balance7)
	return q.LotQuantity
}

func (q *QuantityBreakdown28) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown28) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown28) AddTypeOfPrice() *TypeOfPrice29Choice {
	q.TypeOfPrice = new(TypeOfPrice29Choice)
	return q.TypeOfPrice
}

func (q *QuantityBreakdown28) AddAccountBaseCurrencyAmounts() *BalanceAmounts2 {
	q.AccountBaseCurrencyAmounts = new(BalanceAmounts2)
	return q.AccountBaseCurrencyAmounts
}

func (q *QuantityBreakdown28) AddInstrumentCurrencyAmounts() *BalanceAmounts2 {
	q.InstrumentCurrencyAmounts = new(BalanceAmounts2)
	return q.InstrumentCurrencyAmounts
}

func (q *QuantityBreakdown28) AddAlternateReportingCurrencyAmounts() *BalanceAmounts2 {
	q.AlternateReportingCurrencyAmounts = new(BalanceAmounts2)
	return q.AlternateReportingCurrencyAmounts
}
