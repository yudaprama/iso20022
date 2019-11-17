package model

// Reason for a closure pending status.
type ClosurePendingStatusReason1 struct {

	// Reason for the closure pending status.
	Code *ClosurePendingStatusReason2Choice `xml:"Cd"`

	// Additional information about the reason for the closure pending status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *ClosurePendingStatusReason1) AddCode() *ClosurePendingStatusReason2Choice {
	c.Code = new(ClosurePendingStatusReason2Choice)
	return c.Code
}

func (c *ClosurePendingStatusReason1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
