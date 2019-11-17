package model

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace2 struct {

	// Total quantity of financial instrument for the referenced holding.
	AggregateBalance *BalanceQuantity1Choice `xml:"AggtBal"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	HoldingValue []*ActiveOrHistoricCurrencyAndAmount `xml:"HldgVal"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	BookValue *ActiveOrHistoricCurrencyAndAmount `xml:"BookVal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation1 `xml:"PricDtls"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms3 `xml:"FrgnXchgDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation1 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace2) AddAggregateBalance() *BalanceQuantity1Choice {
	a.AggregateBalance = new(BalanceQuantity1Choice)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace2) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace2) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalancePerSafekeepingPlace2) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace2) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace2) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace2) AddPriceDetails() *PriceInformation1 {
	newValue := new(PriceInformation1)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace2) AddForeignExchangeDetails() *ForeignExchangeTerms3 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms3)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalancePerSafekeepingPlace2) AddBalanceBreakdownDetails() *SubBalanceInformation1 {
	newValue := new(SubBalanceInformation1)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace2) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation {
	newValue := new(AdditionalBalanceInformation)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
