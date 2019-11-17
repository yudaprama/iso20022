package model

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation30 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument other than a investment funds.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Elements used to calculate the valuation haircut.
	ValuationHaircutDetails *BasicCollateralValuation1Details `xml:"ValtnHrcutDtls,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance6 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *Balance8 `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity8Choice `xml:"NotAvlblBal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

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

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation15 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation15 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace28 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation30) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation30) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation30) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation30) AddValuationHaircutDetails() *BasicCollateralValuation1Details {
	a.ValuationHaircutDetails = new(BasicCollateralValuation1Details)
	return a.ValuationHaircutDetails
}

func (a *AggregateBalanceInformation30) AddAggregateBalance() *Balance6 {
	a.AggregateBalance = new(Balance6)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation30) AddAvailableBalance() *Balance8 {
	a.AvailableBalance = new(Balance8)
	return a.AvailableBalance
}

func (a *AggregateBalanceInformation30) AddNotAvailableBalance() *BalanceQuantity8Choice {
	a.NotAvailableBalance = new(BalanceQuantity8Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalanceInformation30) AddSafekeepingPlace() *SafeKeepingPlace1 {
	a.SafekeepingPlace = new(SafeKeepingPlace1)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation30) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation30) AddPriceDetails() *PriceInformation12 {
	newValue := new(PriceInformation12)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation30) AddForeignExchangeDetails() *ForeignExchangeTerms22 {
	newValue := new(ForeignExchangeTerms22)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation30) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation30) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation30) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation30) AddQuantityBreakdown() *QuantityBreakdown27 {
	newValue := new(QuantityBreakdown27)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation30) AddBalanceBreakdown() *SubBalanceInformation15 {
	newValue := new(SubBalanceInformation15)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation30) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation15 {
	newValue := new(AdditionalBalanceInformation15)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation30) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace28 {
	newValue := new(AggregateBalancePerSafekeepingPlace28)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation30) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation30) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
