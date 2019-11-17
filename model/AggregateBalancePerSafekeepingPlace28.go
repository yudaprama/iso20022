package model

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace28 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specify the entity to which the financial instruments are pledged.
	Pledgee *Pledgee1 `xml:"Pldgee,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance6 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *Balance8 `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity8Choice `xml:"NotAvlblBal,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation12 `xml:"PricDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms22 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts3 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts3 `xml:"InstrmCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown27 `xml:"QtyBrkdwn,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation15 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation15 `xml:"AddtlBalBrkdwn,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace28) AddSafekeepingPlace() *SafeKeepingPlace1 {
	a.SafekeepingPlace = new(SafeKeepingPlace1)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace28) AddPlaceOfListing() *MarketIdentification3Choice {
	a.PlaceOfListing = new(MarketIdentification3Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace28) AddPledgee() *Pledgee1 {
	a.Pledgee = new(Pledgee1)
	return a.Pledgee
}

func (a *AggregateBalancePerSafekeepingPlace28) AddAggregateBalance() *Balance6 {
	a.AggregateBalance = new(Balance6)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace28) AddAvailableBalance() *Balance8 {
	a.AvailableBalance = new(Balance8)
	return a.AvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace28) AddNotAvailableBalance() *BalanceQuantity8Choice {
	a.NotAvailableBalance = new(BalanceQuantity8Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace28) AddPriceDetails() *PriceInformation12 {
	newValue := new(PriceInformation12)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace28) AddForeignExchangeDetails() *ForeignExchangeTerms22 {
	newValue := new(ForeignExchangeTerms22)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace28) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace28) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace28) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace28) AddQuantityBreakdown() *QuantityBreakdown27 {
	newValue := new(QuantityBreakdown27)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace28) AddExposureType() *ExposureType16Choice {
	a.ExposureType = new(ExposureType16Choice)
	return a.ExposureType
}

func (a *AggregateBalancePerSafekeepingPlace28) AddBalanceBreakdown() *SubBalanceInformation15 {
	newValue := new(SubBalanceInformation15)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace28) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation15 {
	newValue := new(AdditionalBalanceInformation15)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace28) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}
