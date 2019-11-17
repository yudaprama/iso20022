package model

// Reason for a closed status.
type ClosedStatusReason1 struct {

	// Reason for the closed account status.
	Code *ClosedStatusReason2Choice `xml:"Cd"`

	// Additional information about the reason for the closed account status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *ClosedStatusReason1) AddCode() *ClosedStatusReason2Choice {
	c.Code = new(ClosedStatusReason2Choice)
	return c.Code
}

func (c *ClosedStatusReason1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
