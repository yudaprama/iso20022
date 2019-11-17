package model

// Provides the elements related to the interest amount calculation.
type InterestAmount1 struct {

	// Indicates whether the interest request is new or updated.
	InterestRequestSequence *InterestRequestSequence1Code `xml:"IntrstReqSeq"`

	// Period for which the calculation has been performed.
	InterestPeriod *DatePeriodDetails `xml:"IntrstPrd"`

	// Amount of money representing an interest payment.
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt"`

	// Agreed date for the interest payment.
	ValueDate *ISODate `xml:"ValDt"`

	// Indicates whether the interest will be settled in cash or rolled in the existing collateral balance.
	InterestMethod *InterestMethod1Code `xml:"IntrstMtd"`

	// Percentage charged for the use of an amount of money, usually expressed at an annual rate. The interest rate is the ratio of the amount of interest paid during a certain period of time compared to the principal amount of the interest bearing financial instrument.
	InterestRate *InterestRate1Choice `xml:"IntrstRate,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethod2Code `xml:"DayCntBsis,omitempty"`

	// Amount or percentage of a cash distribution that will be withheld by a tax authority.
	AppliedWithholdingTax *YesNoIndicator `xml:"ApldWhldgTax,omitempty"`

	// Specifies whether the interest is simple or compounded.
	CalculationMethod *CalculationMethod1Code `xml:"ClctnMtd,omitempty"`

	// Specifies the periodicity of the calculation of the interest.
	CalculationFrequency *Frequency1Code `xml:"ClctnFrqcy,omitempty"`

	// Specifies whether the collateral has been posted against the variation margin, the segregated independent amount or to cover any other risk defined with a proprietary code.
	CollateralPurpose *CollateralPurpose1Choice `xml:"CollPurp"`

	// Provides details about the opening collateral balance.
	OpeningCollateralBalance *CollateralBalance1 `xml:"OpngCollBal,omitempty"`

	// Provides details about the closing collateral balance.
	ClosingCollateralBalance *CollateralBalance1 `xml:"ClsgCollBal"`

	// Identifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Additionnal information related to interest request.
	AdditionalInformation *Max210Text `xml:"AddtlInf,omitempty"`

	// Additional references linked to the updated interest payement request.
	ReferenceDetails *Reference20 `xml:"RefDtls,omitempty"`
}

func (i *InterestAmount1) SetInterestRequestSequence(value string) {
	i.InterestRequestSequence = (*InterestRequestSequence1Code)(&value)
}

func (i *InterestAmount1) AddInterestPeriod() *DatePeriodDetails {
	i.InterestPeriod = new(DatePeriodDetails)
	return i.InterestPeriod
}

func (i *InterestAmount1) SetAccruedInterestAmount(value, currency string) {
	i.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InterestAmount1) SetValueDate(value string) {
	i.ValueDate = (*ISODate)(&value)
}

func (i *InterestAmount1) SetInterestMethod(value string) {
	i.InterestMethod = (*InterestMethod1Code)(&value)
}

func (i *InterestAmount1) AddInterestRate() *InterestRate1Choice {
	i.InterestRate = new(InterestRate1Choice)
	return i.InterestRate
}

func (i *InterestAmount1) SetDayCountBasis(value string) {
	i.DayCountBasis = (*InterestComputationMethod2Code)(&value)
}

func (i *InterestAmount1) SetAppliedWithholdingTax(value string) {
	i.AppliedWithholdingTax = (*YesNoIndicator)(&value)
}

func (i *InterestAmount1) SetCalculationMethod(value string) {
	i.CalculationMethod = (*CalculationMethod1Code)(&value)
}

func (i *InterestAmount1) SetCalculationFrequency(value string) {
	i.CalculationFrequency = (*Frequency1Code)(&value)
}

func (i *InterestAmount1) AddCollateralPurpose() *CollateralPurpose1Choice {
	i.CollateralPurpose = new(CollateralPurpose1Choice)
	return i.CollateralPurpose
}

func (i *InterestAmount1) AddOpeningCollateralBalance() *CollateralBalance1 {
	i.OpeningCollateralBalance = new(CollateralBalance1)
	return i.OpeningCollateralBalance
}

func (i *InterestAmount1) AddClosingCollateralBalance() *CollateralBalance1 {
	i.ClosingCollateralBalance = new(CollateralBalance1)
	return i.ClosingCollateralBalance
}

func (i *InterestAmount1) SetStandardSettlementInstructions(value string) {
	i.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (i *InterestAmount1) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max210Text)(&value)
}

func (i *InterestAmount1) AddReferenceDetails() *Reference20 {
	i.ReferenceDetails = new(Reference20)
	return i.ReferenceDetails
}
