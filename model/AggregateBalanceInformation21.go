package model

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation21 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument other than a investment funds.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Elements used to calculate the valuation haircut.
	ValuationHaircutDetails *BasicCollateralValuation1Details `xml:"ValtnHrcutDtls,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *BalanceQuantity5Choice `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity5Choice `xml:"NotAvlblBal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation5 `xml:"PricDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms14 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts3 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts3 `xml:"InstrmCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown23 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation11 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation11 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace20 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation21) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation21) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation21) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation21) AddValuationHaircutDetails() *BasicCollateralValuation1Details {
	a.ValuationHaircutDetails = new(BasicCollateralValuation1Details)
	return a.ValuationHaircutDetails
}

func (a *AggregateBalanceInformation21) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation21) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalanceInformation21) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalanceInformation21) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation21) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation21) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation21) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation21) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation21) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation21) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation21) AddQuantityBreakdown() *QuantityBreakdown23 {
	newValue := new(QuantityBreakdown23)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation21) AddBalanceBreakdown() *SubBalanceInformation11 {
	newValue := new(SubBalanceInformation11)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation21) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation11 {
	newValue := new(AdditionalBalanceInformation11)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation21) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace20 {
	newValue := new(AggregateBalancePerSafekeepingPlace20)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation21) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation21) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
