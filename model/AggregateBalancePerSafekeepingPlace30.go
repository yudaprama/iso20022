package model

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace30 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Choice between formats for the entity to which the financial instruments are pledged.
	Pledgee *Pledgee2 `xml:"Pldgee,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance10 `xml:"AggtBal"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation14 `xml:"PricDtls"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms31 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts5 `xml:"AcctBaseCcyAmts"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts5 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts5 `xml:"AltrnRptgCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown39 `xml:"QtyBrkdwn,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation16 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation16 `xml:"AddtlBalBrkdwn,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *RestrictedFINXMax350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace30) AddSafekeepingPlace() *SafeKeepingPlace2 {
	a.SafekeepingPlace = new(SafeKeepingPlace2)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace30) AddPlaceOfListing() *MarketIdentification4Choice {
	a.PlaceOfListing = new(MarketIdentification4Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace30) AddPledgee() *Pledgee2 {
	a.Pledgee = new(Pledgee2)
	return a.Pledgee
}

func (a *AggregateBalancePerSafekeepingPlace30) AddAggregateBalance() *Balance10 {
	a.AggregateBalance = new(Balance10)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace30) AddPriceDetails() *PriceInformation14 {
	newValue := new(PriceInformation14)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace30) AddForeignExchangeDetails() *ForeignExchangeTerms31 {
	newValue := new(ForeignExchangeTerms31)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace30) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace30) AddAccountBaseCurrencyAmounts() *BalanceAmounts5 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts5)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace30) AddInstrumentCurrencyAmounts() *BalanceAmounts5 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts5)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace30) AddAlternateReportingCurrencyAmounts() *BalanceAmounts5 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts5)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace30) AddQuantityBreakdown() *QuantityBreakdown39 {
	newValue := new(QuantityBreakdown39)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace30) AddExposureType() *ExposureType17Choice {
	a.ExposureType = new(ExposureType17Choice)
	return a.ExposureType
}

func (a *AggregateBalancePerSafekeepingPlace30) AddBalanceBreakdown() *SubBalanceInformation16 {
	newValue := new(SubBalanceInformation16)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace30) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation16 {
	newValue := new(AdditionalBalanceInformation16)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace30) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}
