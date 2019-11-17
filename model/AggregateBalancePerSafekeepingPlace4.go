package model

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace4 struct {

	// Total quantity of financial instrument for the referenced holding.
	AggregateBalance *BalanceQuantity1Choice `xml:"AggtBal"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	HoldingValue []*ActiveOrHistoricCurrencyAndAmount `xml:"HldgVal"`

	// Previous total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	PreviousHoldingValue *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsHldgVal,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Indicates whether the accrued interest is a positive or negative amount.
	AccruedInterestAmountSign *PlusOrMinusIndicator `xml:"AcrdIntrstAmtSgn,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	BookValue *ActiveOrHistoricCurrencyAndAmount `xml:"BookVal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation2 `xml:"PricDtls"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation2 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation2 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace4) AddAggregateBalance() *BalanceQuantity1Choice {
	a.AggregateBalance = new(BalanceQuantity1Choice)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace4) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace4) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalancePerSafekeepingPlace4) SetPreviousHoldingValue(value, currency string) {
	a.PreviousHoldingValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace4) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace4) SetAccruedInterestAmountSign(value string) {
	a.AccruedInterestAmountSign = (*PlusOrMinusIndicator)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace4) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace4) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace4) AddPriceDetails() *PriceInformation2 {
	newValue := new(PriceInformation2)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace4) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalancePerSafekeepingPlace4) AddBalanceBreakdownDetails() *SubBalanceInformation2 {
	newValue := new(SubBalanceInformation2)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace4) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation2 {
	newValue := new(AdditionalBalanceInformation2)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
