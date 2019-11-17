package model

// Provides information related to the collateral position, that is, the identification of the exposed party, the total exposure amount and the total collateral amount held by the taker. It also contains the valuation dates and the requested settlement date of collateral if there is a shortage of collateral.
type Summary1 struct {

	// Sum of the exposures of all transactions which are in the favour of party A. That is, all transactions which would have an amount payable by party B to party A if they were being terminated.
	ExposedAmountPartyA *ActiveCurrencyAndAmount `xml:"XpsdAmtPtyA,omitempty"`

	// Sum of the exposures of all transactions which are in the favour of party B. That is, all transactions which would have an amount payable by party A to party B if they were being terminated.
	ExposedAmountPartyB *ActiveCurrencyAndAmount `xml:"XpsdAmtPtyB,omitempty"`

	// Specifies the underlying business area/type of trade that triggered the posting of collateral.
	ExposureType *ExposureType1Code `xml:"XpsrTp"`

	// Total value of the collateral (post-haircut) held by the exposed party
	TotalValueOfCollateral *ActiveCurrencyAndAmount `xml:"TtlValOfColl"`

	// Specifies the amount of collateral in excess or deficit compared to the exposure.
	NetExcessDeficit *ActiveCurrencyAndAmount `xml:"NetXcssDfcit,omitempty"`

	// Indicates whether the collateral actually posted is a long or a short position.
	NetExcessDeficitIndicator *ShortLong1Code `xml:"NetXcssDfcitInd,omitempty"`

	// Date/time at which the collateral was valued.
	ValuationDateTime *ISODateTime `xml:"ValtnDtTm"`

	// Date on which the instructing party requests settlement of the collateral to take place.
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Provides the more details on the valuation of the collateral that is posted.
	SummaryDetails *SummaryAmounts1 `xml:"SummryDtls,omitempty"`
}

func (s *Summary1) SetExposedAmountPartyA(value, currency string) {
	s.ExposedAmountPartyA = NewActiveCurrencyAndAmount(value, currency)
}

func (s *Summary1) SetExposedAmountPartyB(value, currency string) {
	s.ExposedAmountPartyB = NewActiveCurrencyAndAmount(value, currency)
}

func (s *Summary1) SetExposureType(value string) {
	s.ExposureType = (*ExposureType1Code)(&value)
}

func (s *Summary1) SetTotalValueOfCollateral(value, currency string) {
	s.TotalValueOfCollateral = NewActiveCurrencyAndAmount(value, currency)
}

func (s *Summary1) SetNetExcessDeficit(value, currency string) {
	s.NetExcessDeficit = NewActiveCurrencyAndAmount(value, currency)
}

func (s *Summary1) SetNetExcessDeficitIndicator(value string) {
	s.NetExcessDeficitIndicator = (*ShortLong1Code)(&value)
}

func (s *Summary1) SetValuationDateTime(value string) {
	s.ValuationDateTime = (*ISODateTime)(&value)
}

func (s *Summary1) SetRequestedSettlementDate(value string) {
	s.RequestedSettlementDate = (*ISODate)(&value)
}

func (s *Summary1) AddSummaryDetails() *SummaryAmounts1 {
	s.SummaryDetails = new(SummaryAmounts1)
	return s.SummaryDetails
}
