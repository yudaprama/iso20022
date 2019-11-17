package model

// Provides the details on the margin calculation per financial instrument or per currency.
type MarginCalculation2 struct {

	// Provides details about the security identification.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Net total of the transaction exposure of all outstanding deals.
	ExposureAmount *Amount2 `xml:"XpsrAmt,omitempty"`

	// Provides the total margin amount.
	TotalMarginAmount *AmountAndDirection20 `xml:"TtlMrgnAmt"`

	// Provides details on the valuation of the collateral on deposit.
	CollateralOnDeposit []*Collateral6 `xml:"CollOnDpst,omitempty"`

	// Minimum requirement (expressed in the reporting currency) for a participant if their requirement falls below a specific amount set by the central counterparty.
	MinimumRequirementDeposit *ActiveCurrencyAndAmount `xml:"MinRqrmntDpst,omitempty"`

	// Provide details on the margin result taking into consideration the total margin amount and the minimum requirements deposit.
	MarginResult *MarginResult1Choice `xml:"MrgnRslt,omitempty"`

	// Provides margin calculation details such as the initial margin amount, the variation margin amount or other margin type amounts.
	MarginTypeAmount *Margin3 `xml:"MrgnTpAmt,omitempty"`
}

func (m *MarginCalculation2) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	m.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return m.FinancialInstrumentIdentification
}

func (m *MarginCalculation2) AddExposureAmount() *Amount2 {
	m.ExposureAmount = new(Amount2)
	return m.ExposureAmount
}

func (m *MarginCalculation2) AddTotalMarginAmount() *AmountAndDirection20 {
	m.TotalMarginAmount = new(AmountAndDirection20)
	return m.TotalMarginAmount
}

func (m *MarginCalculation2) AddCollateralOnDeposit() *Collateral6 {
	newValue := new(Collateral6)
	m.CollateralOnDeposit = append(m.CollateralOnDeposit, newValue)
	return newValue
}

func (m *MarginCalculation2) SetMinimumRequirementDeposit(value, currency string) {
	m.MinimumRequirementDeposit = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginCalculation2) AddMarginResult() *MarginResult1Choice {
	m.MarginResult = new(MarginResult1Choice)
	return m.MarginResult
}

func (m *MarginCalculation2) AddMarginTypeAmount() *Margin3 {
	m.MarginTypeAmount = new(Margin3)
	return m.MarginTypeAmount
}
