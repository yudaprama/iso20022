package model

// Reason for a cancelled status.
type CancelledStatusReason16 struct {

	// Reason for the cancelled status.
	Reason *CancelledReason12Choice `xml:"Rsn,omitempty"`

	// Additional information about the cancelled status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CancelledStatusReason16) AddReason() *CancelledReason12Choice {
	c.Reason = new(CancelledReason12Choice)
	return c.Reason
}

func (c *CancelledStatusReason16) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
