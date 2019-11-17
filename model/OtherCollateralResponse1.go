package model

// Provides more details on the response such as the response type, the collateral identification, and optionally further details in case of rejection.
type OtherCollateralResponse1 struct {

	// Specifies the status of the collateral proposal.
	ResponseType *Status4Code `xml:"RspnTp"`

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Specifies the reason why the instruction/cancellation request has a rejected status.
	RejectionReason *RejectionReasonV021Code `xml:"RjctnRsn,omitempty"`

	// Additional information regarding why the collateral proposal has a rejected status.
	RejectionInformation *Max35Text `xml:"RjctnInf,omitempty"`
}

func (o *OtherCollateralResponse1) SetResponseType(value string) {
	o.ResponseType = (*Status4Code)(&value)
}

func (o *OtherCollateralResponse1) SetCollateralIdentification(value string) {
	o.CollateralIdentification = (*Max35Text)(&value)
}

func (o *OtherCollateralResponse1) SetAssetNumber(value string) {
	o.AssetNumber = (*Max35Text)(&value)
}

func (o *OtherCollateralResponse1) SetRejectionReason(value string) {
	o.RejectionReason = (*RejectionReasonV021Code)(&value)
}

func (o *OtherCollateralResponse1) SetRejectionInformation(value string) {
	o.RejectionInformation = (*Max35Text)(&value)
}
