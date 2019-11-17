package model

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation31 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Details of the swap contract.
	AdditionalDerivativeAttributes *DerivativeBasicAttributes1 `xml:"AddtlDerivAttrbts,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance6 `xml:"AggtBal"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation12 `xml:"PricDtls"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms22 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts1 `xml:"AcctBaseCcyAmts"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts1 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts1 `xml:"AltrnRptgCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown28 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation14 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation14 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace29 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation31) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation31) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation31) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation31) AddAdditionalDerivativeAttributes() *DerivativeBasicAttributes1 {
	a.AdditionalDerivativeAttributes = new(DerivativeBasicAttributes1)
	return a.AdditionalDerivativeAttributes
}

func (a *AggregateBalanceInformation31) AddAggregateBalance() *Balance6 {
	a.AggregateBalance = new(Balance6)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation31) AddSafekeepingPlace() *SafeKeepingPlace1 {
	a.SafekeepingPlace = new(SafeKeepingPlace1)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation31) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation31) AddPriceDetails() *PriceInformation12 {
	newValue := new(PriceInformation12)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation31) AddForeignExchangeDetails() *ForeignExchangeTerms22 {
	newValue := new(ForeignExchangeTerms22)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation31) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation31) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation31) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation31) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalanceInformation31) AddQuantityBreakdown() *QuantityBreakdown28 {
	newValue := new(QuantityBreakdown28)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation31) AddBalanceBreakdown() *SubBalanceInformation14 {
	newValue := new(SubBalanceInformation14)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation31) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation14 {
	newValue := new(AdditionalBalanceInformation14)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation31) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace29 {
	newValue := new(AggregateBalancePerSafekeepingPlace29)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation31) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation31) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
