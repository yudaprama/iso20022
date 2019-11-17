package model

// Consideration, such as amount of money,  paid or received in exchange for an amount of money that has been invested, loaned or borrowed for a certain period.
type InterestResult1 struct {

	// Amount of money representing an interest payment.
	InterestDueToA *ActiveCurrencyAndAmount `xml:"IntrstDueToA,omitempty"`

	// Amount of money representing an interest payment.
	InterestDueToB *ActiveCurrencyAndAmount `xml:"IntrstDueToB,omitempty"`

	// Agreed date for the interest payment.
	ValueDate *ISODate `xml:"ValDt"`

	// Indicates whether the interest will be settled in cash or rolled in the existing collateral balance.
	InterestMethod *InterestMethod1Code `xml:"IntrstMtd"`

	// Provides details about the opening collateral balance.
	OpeningCollateralBalance *CollateralBalance1 `xml:"OpngCollBal,omitempty"`

	// Provides details about the closing collateral balance.
	ClosingCollateralBalance *CollateralBalance1 `xml:"ClsgCollBal"`
}

func (i *InterestResult1) SetInterestDueToA(value, currency string) {
	i.InterestDueToA = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InterestResult1) SetInterestDueToB(value, currency string) {
	i.InterestDueToB = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InterestResult1) SetValueDate(value string) {
	i.ValueDate = (*ISODate)(&value)
}

func (i *InterestResult1) SetInterestMethod(value string) {
	i.InterestMethod = (*InterestMethod1Code)(&value)
}

func (i *InterestResult1) AddOpeningCollateralBalance() *CollateralBalance1 {
	i.OpeningCollateralBalance = new(CollateralBalance1)
	return i.OpeningCollateralBalance
}

func (i *InterestResult1) AddClosingCollateralBalance() *CollateralBalance1 {
	i.ClosingCollateralBalance = new(CollateralBalance1)
	return i.ClosingCollateralBalance
}
