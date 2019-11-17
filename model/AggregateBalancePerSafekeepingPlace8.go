package model

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace8 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification5 `xml:"PlcOfListg,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *BalanceQuantity5Choice `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity5Choice `xml:"NotAvlblBal,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation5 `xml:"PricDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms1 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts3 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts3 `xml:"InstrmCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation5 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation5 `xml:"AddtlBalBrkdwn,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace8) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace8) AddPlaceOfListing() *MarketIdentification5 {
	a.PlaceOfListing = new(MarketIdentification5)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace8) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace8) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace8) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace8) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace8) AddForeignExchangeDetails() *ForeignExchangeTerms1 {
	newValue := new(ForeignExchangeTerms1)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace8) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace8) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace8) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace8) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace8) AddBalanceBreakdown() *SubBalanceInformation5 {
	newValue := new(SubBalanceInformation5)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace8) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation5 {
	newValue := new(AdditionalBalanceInformation5)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace8) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}
