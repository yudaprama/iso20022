package model

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace31 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specify the entity to which the financial instruments are pledged.
	Pledgee *Pledgee2 `xml:"Pldgee,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance10 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *Balance12 `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity12Choice `xml:"NotAvlblBal,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation14 `xml:"PricDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms31 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts4 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts4 `xml:"InstrmCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown40 `xml:"QtyBrkdwn,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation17 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation17 `xml:"AddtlBalBrkdwn,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *RestrictedFINXMax350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace31) AddSafekeepingPlace() *SafeKeepingPlace2 {
	a.SafekeepingPlace = new(SafeKeepingPlace2)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace31) AddPlaceOfListing() *MarketIdentification4Choice {
	a.PlaceOfListing = new(MarketIdentification4Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace31) AddPledgee() *Pledgee2 {
	a.Pledgee = new(Pledgee2)
	return a.Pledgee
}

func (a *AggregateBalancePerSafekeepingPlace31) AddAggregateBalance() *Balance10 {
	a.AggregateBalance = new(Balance10)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace31) AddAvailableBalance() *Balance12 {
	a.AvailableBalance = new(Balance12)
	return a.AvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace31) AddNotAvailableBalance() *BalanceQuantity12Choice {
	a.NotAvailableBalance = new(BalanceQuantity12Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace31) AddPriceDetails() *PriceInformation14 {
	newValue := new(PriceInformation14)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace31) AddForeignExchangeDetails() *ForeignExchangeTerms31 {
	newValue := new(ForeignExchangeTerms31)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace31) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace31) AddAccountBaseCurrencyAmounts() *BalanceAmounts4 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts4)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace31) AddInstrumentCurrencyAmounts() *BalanceAmounts4 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts4)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace31) AddQuantityBreakdown() *QuantityBreakdown40 {
	newValue := new(QuantityBreakdown40)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace31) AddExposureType() *ExposureType17Choice {
	a.ExposureType = new(ExposureType17Choice)
	return a.ExposureType
}

func (a *AggregateBalancePerSafekeepingPlace31) AddBalanceBreakdown() *SubBalanceInformation17 {
	newValue := new(SubBalanceInformation17)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace31) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation17 {
	newValue := new(AdditionalBalanceInformation17)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace31) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}
