package model

// Details of the margin call request.
type MarginCall1 struct {

	// Sum of the exposures of all transactions which are in the favour of party A. That is, all transactions which would have an amount payable by party B to party A if they were being terminated.
	ExposedAmountPartyA *ActiveCurrencyAndAmount `xml:"XpsdAmtPtyA,omitempty"`

	// Sum of the exposures of all transactions which are in the favour of party B. That is, all transactions which would have an amount payable by party A to party B if they were being terminated.
	ExposedAmountPartyB *ActiveCurrencyAndAmount `xml:"XpsdAmtPtyB,omitempty"`

	// Determines how the variation margin requirement is to be calculated:
	// - either Net, in which the exposure of all transactions in favour of party A and the the exposure of all transactions in favour of party B will be netted together or
	// - gross in which two separate variation margin requirements will be determined.
	ExposureConvention *ExposureConventionType1Code `xml:"XpsrCnvntn,omitempty"`

	// Amount applied as an add-on to the exposure (to party A) usually intended to cover a possible increase in exposure before the next valuation date.
	IndependentAmountPartyA *AggregatedIndependentAmount1 `xml:"IndpdntAmtPtyA,omitempty"`

	// An amount applied as an add-on to the exposure (to party B) usually intended to cover a possible increase in exposure before the next valuation date.
	IndependentAmountPartyB *AggregatedIndependentAmount1 `xml:"IndpdntAmtPtyB,omitempty"`

	// Provides information like threshold amount, threshold type, minimum transfer amount, rouding amount or rounding convention, that applies to either the variation margin or the segregated independent amount.
	MarginTerms *MarginTerms1Choice `xml:"MrgnTerms,omitempty"`

	// Provides details about the collateral held, in transit or that still needs to be agreed by both parties with a segregation between variation margin and segregated independent amount.
	CollateralBalance *CollateralBalance1Choice `xml:"CollBal,omitempty"`
}

func (m *MarginCall1) SetExposedAmountPartyA(value, currency string) {
	m.ExposedAmountPartyA = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginCall1) SetExposedAmountPartyB(value, currency string) {
	m.ExposedAmountPartyB = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginCall1) SetExposureConvention(value string) {
	m.ExposureConvention = (*ExposureConventionType1Code)(&value)
}

func (m *MarginCall1) AddIndependentAmountPartyA() *AggregatedIndependentAmount1 {
	m.IndependentAmountPartyA = new(AggregatedIndependentAmount1)
	return m.IndependentAmountPartyA
}

func (m *MarginCall1) AddIndependentAmountPartyB() *AggregatedIndependentAmount1 {
	m.IndependentAmountPartyB = new(AggregatedIndependentAmount1)
	return m.IndependentAmountPartyB
}

func (m *MarginCall1) AddMarginTerms() *MarginTerms1Choice {
	m.MarginTerms = new(MarginTerms1Choice)
	return m.MarginTerms
}

func (m *MarginCall1) AddCollateralBalance() *CollateralBalance1Choice {
	m.CollateralBalance = new(CollateralBalance1Choice)
	return m.CollateralBalance
}
