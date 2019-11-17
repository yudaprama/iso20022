package model

// Provides the total margin amount, the collateral amount on deposit and the total minimum requirement that used to calculate the margin result, either an excess or a deficit.
type MarginCalculation1 struct {

	// Total margin requirement (expressed in the reporting currency) that must be provided by the clearing member to the central counterparty. This is the total requirement calculated to cover the initial margin and the variation margin.
	TotalMarginAmount *AmountAndDirection20 `xml:"TtlMrgnAmt"`

	// Provides details on the valuation of the collateral on deposit.
	CollateralOnDeposit []*Collateral6 `xml:"CollOnDpst,omitempty"`

	// Minimum requirement (expressed in the reporting currency) for a participant if their requirement falls below a specific amount set by the central counterparty.
	MinimumRequirementDeposit *ActiveCurrencyAndAmount `xml:"MinRqrmntDpst,omitempty"`

	// Provide details on the margin result taking into consideration the total margin amount and the minimum requirements deposit.
	MarginResult *MarginResult1Choice `xml:"MrgnRslt,omitempty"`
}

func (m *MarginCalculation1) AddTotalMarginAmount() *AmountAndDirection20 {
	m.TotalMarginAmount = new(AmountAndDirection20)
	return m.TotalMarginAmount
}

func (m *MarginCalculation1) AddCollateralOnDeposit() *Collateral6 {
	newValue := new(Collateral6)
	m.CollateralOnDeposit = append(m.CollateralOnDeposit, newValue)
	return newValue
}

func (m *MarginCalculation1) SetMinimumRequirementDeposit(value, currency string) {
	m.MinimumRequirementDeposit = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginCalculation1) AddMarginResult() *MarginResult1Choice {
	m.MarginResult = new(MarginResult1Choice)
	return m.MarginResult
}
