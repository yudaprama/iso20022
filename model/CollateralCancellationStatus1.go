package model

// Provides details on the status (that is accept or reject) of the CollateralManagementCancellationRequest message.
type CollateralCancellationStatus1 struct {

	// Allows to provide a cancellation status using a code or a proprietary status.
	CollateralStatusCode *Status4Code `xml:"CollStsCd"`

	// Provides additional information on the status of the CollateralManagementCancellationRequest message.
	AdditionalInformation *Max35Text `xml:"AddtlInf,omitempty"`

	// Provides rejection reason and optionaly additional information.
	RejectionDetails *RejectionStatus2 `xml:"RjctnDtls,omitempty"`
}

func (c *CollateralCancellationStatus1) SetCollateralStatusCode(value string) {
	c.CollateralStatusCode = (*Status4Code)(&value)
}

func (c *CollateralCancellationStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max35Text)(&value)
}

func (c *CollateralCancellationStatus1) AddRejectionDetails() *RejectionStatus2 {
	c.RejectionDetails = new(RejectionStatus2)
	return c.RejectionDetails
}
