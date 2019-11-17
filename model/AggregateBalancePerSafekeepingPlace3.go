package model

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace3 struct {

	// Total quantity of financial instrument for the referenced holding.
	AggregateQuantity *BalanceQuantity1Choice `xml:"AggtQty"`

	// Total quantity of financial instrument for the referenced holding that is available.
	AvailableQuantity *BalanceQuantity1Choice `xml:"AvlblQty,omitempty"`

	// Total quantity of financial instrument for the referenced holding that is not available.
	NotAvailableQuantity *BalanceQuantity1Choice `xml:"NotAvlblQty,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	HoldingValue []*ActiveOrHistoricCurrencyAndAmount `xml:"HldgVal,omitempty"`

	// Previous total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	PreviousHoldingValue *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsHldgVal,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Indicates whether the accrued interest is a positive or negative amount.
	AccruedInterestAmountSign *PlusOrMinusIndicator `xml:"AcrdIntrstAmtSgn,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	BookValue *ActiveOrHistoricCurrencyAndAmount `xml:"BookVal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation2 `xml:"PricDtls,omitempty"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation2 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation2 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace3) AddAggregateQuantity() *BalanceQuantity1Choice {
	a.AggregateQuantity = new(BalanceQuantity1Choice)
	return a.AggregateQuantity
}

func (a *AggregateBalancePerSafekeepingPlace3) AddAvailableQuantity() *BalanceQuantity1Choice {
	a.AvailableQuantity = new(BalanceQuantity1Choice)
	return a.AvailableQuantity
}

func (a *AggregateBalancePerSafekeepingPlace3) AddNotAvailableQuantity() *BalanceQuantity1Choice {
	a.NotAvailableQuantity = new(BalanceQuantity1Choice)
	return a.NotAvailableQuantity
}

func (a *AggregateBalancePerSafekeepingPlace3) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace3) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalancePerSafekeepingPlace3) SetPreviousHoldingValue(value, currency string) {
	a.PreviousHoldingValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace3) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace3) SetAccruedInterestAmountSign(value string) {
	a.AccruedInterestAmountSign = (*PlusOrMinusIndicator)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace3) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace3) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace3) AddPriceDetails() *PriceInformation2 {
	newValue := new(PriceInformation2)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace3) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalancePerSafekeepingPlace3) AddBalanceBreakdownDetails() *SubBalanceInformation2 {
	newValue := new(SubBalanceInformation2)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace3) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation2 {
	newValue := new(AdditionalBalanceInformation2)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
