package model

// Provides more details on the response such as the response type, the collateral identification, and optionally further details in case of rejection.
type SecuritiesCollateralResponse1 struct {

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Specifies the status of the collateral proposal.
	ResponseType *Status4Code `xml:"RspnTp"`

	// Specifies the reason why the instruction/cancellation request has a rejected status.
	RejectionReason *RejectionReasonV021Code `xml:"RjctnRsn,omitempty"`

	// Additional information regarding why the collateral proposal has a rejected status.
	RejectionInformation *Max35Text `xml:"RjctnInf,omitempty"`
}

func (s *SecuritiesCollateralResponse1) SetCollateralIdentification(value string) {
	s.CollateralIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesCollateralResponse1) SetAssetNumber(value string) {
	s.AssetNumber = (*Max35Text)(&value)
}

func (s *SecuritiesCollateralResponse1) SetResponseType(value string) {
	s.ResponseType = (*Status4Code)(&value)
}

func (s *SecuritiesCollateralResponse1) SetRejectionReason(value string) {
	s.RejectionReason = (*RejectionReasonV021Code)(&value)
}

func (s *SecuritiesCollateralResponse1) SetRejectionInformation(value string) {
	s.RejectionInformation = (*Max35Text)(&value)
}
