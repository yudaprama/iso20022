package model

// Provides amounts taken in to account to calculate the collateral position.
type SummaryAmounts1 struct {

	// Amount of unsecured exposure a counterparty will accept before issuing a margin call in the base currency.
	ThresholdAmount *ActiveCurrencyAndAmount `xml:"ThrshldAmt,omitempty"`

	// Specifies if the threshold amount is secured or unsecured.
	ThresholdType *ThresholdType1Code `xml:"ThrshldTp,omitempty"`

	// Total value of posted collateral (pre-haircut) held by the taker.
	PreHaircutCollateralValue *ActiveCurrencyAndAmount `xml:"PreHrcutCollVal,omitempty"`

	// Total amount of collateral required (unrounded).
	AdjustedExposure *ActiveCurrencyAndAmount `xml:"AdjstdXpsr,omitempty"`

	// Total amount of collateral required (rounded).
	CollateralRequired *ActiveCurrencyAndAmount `xml:"CollReqrd,omitempty"`

	// Minimum amount to pay/receive as specified in the agreement in the base currency (to avoid the need to transfer an inconveniently small amount of collateral).
	MinimumTransferAmount *ActiveCurrencyAndAmount `xml:"MinTrfAmt,omitempty"`

	// Amount specified to avoid the need to transfer uneven amounts of collateral.
	RoundingAmount *ActiveCurrencyAndAmount `xml:"RndgAmt,omitempty"`

	// Exposure value at previous valuation.
	PreviousExposureValue *ActiveCurrencyAndAmount `xml:"PrvsXpsrVal,omitempty"`

	// Value of collateral at previous valuation.
	PreviousCollateralValue *ActiveCurrencyAndAmount `xml:"PrvsCollVal,omitempty"`

	// Value of incoming collateral, to be settled.
	TotalPendingIncomingCollateral *ActiveCurrencyAndAmount `xml:"TtlPdgIncmgColl,omitempty"`

	// Value of outgoing collateral, to be settled.
	TotalPendingOutgoingCollateral *ActiveCurrencyAndAmount `xml:"TtlPdgOutgngColl,omitempty"`

	// Sum of accrued interest.
	TotalAccruedInterestAmount *ActiveCurrencyAndAmount `xml:"TtlAcrdIntrstAmt,omitempty"`

	// Sum of fees/commissions.
	TotalFees *ActiveCurrencyAndAmount `xml:"TtlFees,omitempty"`
}

func (s *SummaryAmounts1) SetThresholdAmount(value, currency string) {
	s.ThresholdAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SummaryAmounts1) SetThresholdType(value string) {
	s.ThresholdType = (*ThresholdType1Code)(&value)
}

func (s *SummaryAmounts1) SetPreHaircutCollateralValue(value, currency string) {
	s.PreHaircutCollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SummaryAmounts1) SetAdjustedExposure(value, currency string) {
	s.AdjustedExposure = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SummaryAmounts1) SetCollateralRequired(value, currency string) {
	s.CollateralRequired = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SummaryAmounts1) SetMinimumTransferAmount(value, currency string) {
	s.MinimumTransferAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SummaryAmounts1) SetRoundingAmount(value, currency string) {
	s.RoundingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SummaryAmounts1) SetPreviousExposureValue(value, currency string) {
	s.PreviousExposureValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SummaryAmounts1) SetPreviousCollateralValue(value, currency string) {
	s.PreviousCollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SummaryAmounts1) SetTotalPendingIncomingCollateral(value, currency string) {
	s.TotalPendingIncomingCollateral = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SummaryAmounts1) SetTotalPendingOutgoingCollateral(value, currency string) {
	s.TotalPendingOutgoingCollateral = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SummaryAmounts1) SetTotalAccruedInterestAmount(value, currency string) {
	s.TotalAccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SummaryAmounts1) SetTotalFees(value, currency string) {
	s.TotalFees = NewActiveCurrencyAndAmount(value, currency)
}
