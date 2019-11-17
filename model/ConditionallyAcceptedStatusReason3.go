package model

// Reason for a conditionally accepted status.
type ConditionallyAcceptedStatusReason3 struct {

	// Reason for the conditionally accepted status expressed as a code.
	Reason *ConditionallyAcceptedStatusReason3Choice `xml:"Rsn"`

	// Additional information about the conditionally accepted reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *ConditionallyAcceptedStatusReason3) AddReason() *ConditionallyAcceptedStatusReason3Choice {
	c.Reason = new(ConditionallyAcceptedStatusReason3Choice)
	return c.Reason
}

func (c *ConditionallyAcceptedStatusReason3) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
