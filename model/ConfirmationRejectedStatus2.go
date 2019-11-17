package model

// Rejection of a confirmation.
type ConfirmationRejectedStatus2 struct {

	// Reason for the rejected status.
	Reason *ConfirmationRejectedReason1Choice `xml:"Rsn,omitempty"`

	// Additional information about the rejected reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *ConfirmationRejectedStatus2) AddReason() *ConfirmationRejectedReason1Choice {
	c.Reason = new(ConfirmationRejectedReason1Choice)
	return c.Reason
}

func (c *ConfirmationRejectedStatus2) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
