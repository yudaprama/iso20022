package model

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation32 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument22 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Details of the swap contract.
	AdditionalDerivativeAttributes *DerivativeBasicAttributes2 `xml:"AddtlDerivAttrbts,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance10 `xml:"AggtBal"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

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

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation16 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation16 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace30 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *RestrictedFINXMax350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation32) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation32) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation32) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument22 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument22)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation32) AddAdditionalDerivativeAttributes() *DerivativeBasicAttributes2 {
	a.AdditionalDerivativeAttributes = new(DerivativeBasicAttributes2)
	return a.AdditionalDerivativeAttributes
}

func (a *AggregateBalanceInformation32) AddAggregateBalance() *Balance10 {
	a.AggregateBalance = new(Balance10)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation32) AddSafekeepingPlace() *SafeKeepingPlace2 {
	a.SafekeepingPlace = new(SafeKeepingPlace2)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation32) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation32) AddPriceDetails() *PriceInformation14 {
	newValue := new(PriceInformation14)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation32) AddForeignExchangeDetails() *ForeignExchangeTerms31 {
	newValue := new(ForeignExchangeTerms31)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation32) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation32) AddAccountBaseCurrencyAmounts() *BalanceAmounts5 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts5)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation32) AddInstrumentCurrencyAmounts() *BalanceAmounts5 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts5)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation32) AddAlternateReportingCurrencyAmounts() *BalanceAmounts5 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts5)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalanceInformation32) AddQuantityBreakdown() *QuantityBreakdown39 {
	newValue := new(QuantityBreakdown39)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation32) AddBalanceBreakdown() *SubBalanceInformation16 {
	newValue := new(SubBalanceInformation16)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation32) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation16 {
	newValue := new(AdditionalBalanceInformation16)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation32) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace30 {
	newValue := new(AggregateBalancePerSafekeepingPlace30)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation32) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (a *AggregateBalanceInformation32) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
