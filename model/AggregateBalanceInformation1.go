package model

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type AggregateBalanceInformation1 struct {

	// Total quantity of financial instrument for the referenced holding.
	AggregateQuantity *BalanceQuantity1Choice `xml:"AggtQty"`

	// Total quantity of financial instrument for the referenced holding that is available.
	AvailableQuantity *BalanceQuantity1Choice `xml:"AvlblQty,omitempty"`

	// Total quantity of financial instrument for the referenced holding that is not available.
	NotAvailableQuantity *BalanceQuantity1Choice `xml:"NotAvlblQty,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	HoldingValue []*ActiveOrHistoricCurrencyAndAmount `xml:"HldgVal,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	BookValue *ActiveOrHistoricCurrencyAndAmount `xml:"BookVal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc,omitempty"`

	// Security held on the account for which the balance is calculated.
	FinancialInstrumentDetails *FinancialInstrument4 `xml:"FinInstrmDtls"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation1 `xml:"PricDtls,omitempty"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms3 `xml:"FrgnXchgDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation1 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation `xml:"AddtlBalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace1 `xml:"BalAtSfkpgPlc,omitempty"`
}

func (a *AggregateBalanceInformation1) AddAggregateQuantity() *BalanceQuantity1Choice {
	a.AggregateQuantity = new(BalanceQuantity1Choice)
	return a.AggregateQuantity
}

func (a *AggregateBalanceInformation1) AddAvailableQuantity() *BalanceQuantity1Choice {
	a.AvailableQuantity = new(BalanceQuantity1Choice)
	return a.AvailableQuantity
}

func (a *AggregateBalanceInformation1) AddNotAvailableQuantity() *BalanceQuantity1Choice {
	a.NotAvailableQuantity = new(BalanceQuantity1Choice)
	return a.NotAvailableQuantity
}

func (a *AggregateBalanceInformation1) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation1) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalanceInformation1) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation1) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation1) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation1) AddFinancialInstrumentDetails() *FinancialInstrument4 {
	a.FinancialInstrumentDetails = new(FinancialInstrument4)
	return a.FinancialInstrumentDetails
}

func (a *AggregateBalanceInformation1) AddPriceDetails() *PriceInformation1 {
	newValue := new(PriceInformation1)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation1) AddForeignExchangeDetails() *ForeignExchangeTerms3 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms3)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalanceInformation1) AddBalanceBreakdownDetails() *SubBalanceInformation1 {
	newValue := new(SubBalanceInformation1)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation1) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation {
	newValue := new(AdditionalBalanceInformation)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation1) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace1 {
	newValue := new(AggregateBalancePerSafekeepingPlace1)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}
