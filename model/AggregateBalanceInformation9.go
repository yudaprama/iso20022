package model

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation9 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument other than a investment funds.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes8 `xml:"FinInstrmAttrbts,omitempty"`

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

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace8 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (a *AggregateBalanceInformation9) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation9) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes8 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes8)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation9) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation9) AddValuationHaircutDetails() *BasicCollateralValuation1Details {
	a.ValuationHaircutDetails = new(BasicCollateralValuation1Details)
	return a.ValuationHaircutDetails
}

func (a *AggregateBalanceInformation9) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation9) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalanceInformation9) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalanceInformation9) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation9) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation9) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation9) AddForeignExchangeDetails() *ForeignExchangeTerms1 {
	newValue := new(ForeignExchangeTerms1)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation9) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation9) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation9) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation9) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation9) AddBalanceBreakdown() *SubBalanceInformation5 {
	newValue := new(SubBalanceInformation5)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation9) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation5 {
	newValue := new(AdditionalBalanceInformation5)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation9) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace8 {
	newValue := new(AggregateBalancePerSafekeepingPlace8)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation9) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation9) AddExtension() *Extension2 {
	newValue := new(Extension2)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
