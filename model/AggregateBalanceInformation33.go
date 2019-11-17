package model

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation33 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument other than a investment funds.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument22 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Elements used to calculate the valuation haircut.
	ValuationHaircutDetails *BasicCollateralValuation1Details `xml:"ValtnHrcutDtls,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance10 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *Balance12 `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity12Choice `xml:"NotAvlblBal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

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

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation17 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation17 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace31 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *RestrictedFINXMax350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation33) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation33) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation33) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument22 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument22)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation33) AddValuationHaircutDetails() *BasicCollateralValuation1Details {
	a.ValuationHaircutDetails = new(BasicCollateralValuation1Details)
	return a.ValuationHaircutDetails
}

func (a *AggregateBalanceInformation33) AddAggregateBalance() *Balance10 {
	a.AggregateBalance = new(Balance10)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation33) AddAvailableBalance() *Balance12 {
	a.AvailableBalance = new(Balance12)
	return a.AvailableBalance
}

func (a *AggregateBalanceInformation33) AddNotAvailableBalance() *BalanceQuantity12Choice {
	a.NotAvailableBalance = new(BalanceQuantity12Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalanceInformation33) AddSafekeepingPlace() *SafeKeepingPlace2 {
	a.SafekeepingPlace = new(SafeKeepingPlace2)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation33) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation33) AddPriceDetails() *PriceInformation14 {
	newValue := new(PriceInformation14)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation33) AddForeignExchangeDetails() *ForeignExchangeTerms31 {
	newValue := new(ForeignExchangeTerms31)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation33) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation33) AddAccountBaseCurrencyAmounts() *BalanceAmounts4 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts4)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation33) AddInstrumentCurrencyAmounts() *BalanceAmounts4 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts4)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation33) AddQuantityBreakdown() *QuantityBreakdown40 {
	newValue := new(QuantityBreakdown40)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation33) AddBalanceBreakdown() *SubBalanceInformation17 {
	newValue := new(SubBalanceInformation17)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation33) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation17 {
	newValue := new(AdditionalBalanceInformation17)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation33) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace31 {
	newValue := new(AggregateBalancePerSafekeepingPlace31)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation33) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (a *AggregateBalanceInformation33) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
