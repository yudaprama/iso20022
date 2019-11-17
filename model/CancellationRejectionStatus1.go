package model

// Status advising on the rejection of the cancellation request.
type CancellationRejectionStatus1 struct {

	// Reason advising the rejection of the instruction cancellation request.
	Reason *RejectionReason2Code `xml:"Rsn"`

	// This code can be used in case another reason is required.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Additional information about the reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CancellationRejectionStatus1) SetReason(value string) {
	c.Reason = (*RejectionReason2Code)(&value)
}

func (c *CancellationRejectionStatus1) SetExtendedReason(value string) {
	c.ExtendedReason = (*Extended350Code)(&value)
}

func (c *CancellationRejectionStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
