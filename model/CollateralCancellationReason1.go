package model

// Reason for which the collateral message has been cancelled.
type CollateralCancellationReason1 struct {

	// Allows to provides additional information on the cancellation reason.
	AdditionalInformation *Max35Text `xml:"AddtlInf,omitempty"`

	// Allows to provide a cancellation reason using a code or proprietary reason.
	CancellationReasonCode *CollateralCancellationType1Choice `xml:"CxlRsnCd"`
}

func (c *CollateralCancellationReason1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max35Text)(&value)
}

func (c *CollateralCancellationReason1) AddCancellationReasonCode() *CollateralCancellationType1Choice {
	c.CancellationReasonCode = new(CollateralCancellationType1Choice)
	return c.CancellationReasonCode
}
