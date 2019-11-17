package model

// Information about the business status of a cancellation request message.
type CancellationStatusInformation1 struct {

	// Information on the business status of the cancellation.
	Status *CancellationStatus4Code `xml:"Sts"`

	// The reason for the cancellation status.
	StatusReason *StatusReason4Choice `xml:"StsRsn,omitempty"`

	// Further details on the cancellation status reason.
	AdditionalStatusReasonInformation []*Max105Text `xml:"AddtlStsRsnInf,omitempty"`
}

func (c *CancellationStatusInformation1) SetStatus(value string) {
	c.Status = (*CancellationStatus4Code)(&value)
}

func (c *CancellationStatusInformation1) AddStatusReason() *StatusReason4Choice {
	c.StatusReason = new(StatusReason4Choice)
	return c.StatusReason
}

func (c *CancellationStatusInformation1) AddAdditionalStatusReasonInformation(value string) {
	c.AdditionalStatusReasonInformation = append(c.AdditionalStatusReasonInformation, (*Max105Text)(&value))
}
