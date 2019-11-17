package model

// Details of breakdown of a quantity.
type QuantityBreakdown4 struct {

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

	// Valuation amounts for the lot provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts2 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts2 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts2 `xml:"AltrnRptgCcyAmts,omitempty"`
}

func (q *QuantityBreakdown4) AddLotNumber() *Number2Choice {
	q.LotNumber = new(Number2Choice)
	return q.LotNumber
}

func (q *QuantityBreakdown4) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown4) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown4) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown4) AddTypeOfPrice() *TypeOfPrice3Choice {
	q.TypeOfPrice = new(TypeOfPrice3Choice)
	return q.TypeOfPrice
}

func (q *QuantityBreakdown4) AddAccountBaseCurrencyAmounts() *BalanceAmounts2 {
	q.AccountBaseCurrencyAmounts = new(BalanceAmounts2)
	return q.AccountBaseCurrencyAmounts
}

func (q *QuantityBreakdown4) AddInstrumentCurrencyAmounts() *BalanceAmounts2 {
	q.InstrumentCurrencyAmounts = new(BalanceAmounts2)
	return q.InstrumentCurrencyAmounts
}

func (q *QuantityBreakdown4) AddAlternateReportingCurrencyAmounts() *BalanceAmounts2 {
	q.AlternateReportingCurrencyAmounts = new(BalanceAmounts2)
	return q.AlternateReportingCurrencyAmounts
}
