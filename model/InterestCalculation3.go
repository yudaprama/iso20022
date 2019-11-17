package model

// Provides the details of the interest calculation.
type InterestCalculation3 struct {

	// Indicates the calculation date of the interest amount.
	CalculationDate *ISODate `xml:"ClctnDt"`

	// Provides the identification of the collateral account.
	CollateralAccountIdentification *AccountIdentification26 `xml:"CollAcctId,omitempty"`

	// Provides the collateral amount used to calculate the interest amount and includes debit/short or credit/long positions.
	EffectivePrincipalAmount *AmountAndDirection20 `xml:"FctvPrncplAmt"`

	// Provides the collateral amount posted before taking into account the collateral movement amount.
	PrincipalAmount *AmountAndDirection20 `xml:"PrncplAmt,omitempty"`

	// Provides the additional amount of collateral posted between two calculation dates.
	MovementAmount *AmountAndDirection20 `xml:"MvmntAmt,omitempty"`

	// Indicates the number of days for the calculation of the interest.
	NumberOfDays *Number `xml:"NbOfDays,omitempty"`

	// Specifies the percentage charged for the use of an amount of money, usually expressed at an annual rate. The interest rate is the ratio of the amount of interest paid during a certain period of time compared to the principal amount of the interest bearing financial instrument.
	EffectiveRate *PercentageRate `xml:"FctvRate"`

	// Specifies the percentage charged for the use of an amount of money, usually expressed at an annual rate. The interest rate is the ratio of the amount of interest paid during a certain period of time compared to the principal amount of the interest bearing financial instrument.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Indicates the differences in interest rates.
	Spread *PercentageRate `xml:"Sprd,omitempty"`

	// Specifies the amount of money representing an interest payment.
	AccruedInterestAmount *AmountAndDirection20 `xml:"AcrdIntrstAmt"`

	// Specifies the total amount of money representing an interest payment.
	AggregatedInterestAmount *ActiveCurrencyAndAmount `xml:"AggtdIntrstAmt,omitempty"`
}

func (i *InterestCalculation3) SetCalculationDate(value string) {
	i.CalculationDate = (*ISODate)(&value)
}

func (i *InterestCalculation3) AddCollateralAccountIdentification() *AccountIdentification26 {
	i.CollateralAccountIdentification = new(AccountIdentification26)
	return i.CollateralAccountIdentification
}

func (i *InterestCalculation3) AddEffectivePrincipalAmount() *AmountAndDirection20 {
	i.EffectivePrincipalAmount = new(AmountAndDirection20)
	return i.EffectivePrincipalAmount
}

func (i *InterestCalculation3) AddPrincipalAmount() *AmountAndDirection20 {
	i.PrincipalAmount = new(AmountAndDirection20)
	return i.PrincipalAmount
}

func (i *InterestCalculation3) AddMovementAmount() *AmountAndDirection20 {
	i.MovementAmount = new(AmountAndDirection20)
	return i.MovementAmount
}

func (i *InterestCalculation3) SetNumberOfDays(value string) {
	i.NumberOfDays = (*Number)(&value)
}

func (i *InterestCalculation3) SetEffectiveRate(value string) {
	i.EffectiveRate = (*PercentageRate)(&value)
}

func (i *InterestCalculation3) SetInterestRate(value string) {
	i.InterestRate = (*PercentageRate)(&value)
}

func (i *InterestCalculation3) SetSpread(value string) {
	i.Spread = (*PercentageRate)(&value)
}

func (i *InterestCalculation3) AddAccruedInterestAmount() *AmountAndDirection20 {
	i.AccruedInterestAmount = new(AmountAndDirection20)
	return i.AccruedInterestAmount
}

func (i *InterestCalculation3) SetAggregatedInterestAmount(value, currency string) {
	i.AggregatedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}
