package model

// Specifies the reason why the corporate action reversal occurs.
type CorporateActionReversalReason3 struct {

	// Specifies the reason for the reversal.
	Reason *CorporateActionReversalReason3Choice `xml:"Rsn"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max256Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CorporateActionReversalReason3) AddReason() *CorporateActionReversalReason3Choice {
	c.Reason = new(CorporateActionReversalReason3Choice)
	return c.Reason
}

func (c *CorporateActionReversalReason3) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max256Text)(&value)
}
