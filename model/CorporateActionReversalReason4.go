package model

// Specifies the reason why the corporate action reversal occurs.
type CorporateActionReversalReason4 struct {

	// Specifies the reason for the reversal.
	Reason *CorporateActionReversalReason4Choice `xml:"Rsn"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax256Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CorporateActionReversalReason4) AddReason() *CorporateActionReversalReason4Choice {
	c.Reason = new(CorporateActionReversalReason4Choice)
	return c.Reason
}

func (c *CorporateActionReversalReason4) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax256Text)(&value)
}
