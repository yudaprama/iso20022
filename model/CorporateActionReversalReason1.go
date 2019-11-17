package model

// Specifies the reason why the corporate action reversal occurs.
type CorporateActionReversalReason1 struct {

	// Specifies the reason for the reversal.
	Reason *CorporateActionReversalReason1Choice `xml:"Rsn"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max256Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CorporateActionReversalReason1) AddReason() *CorporateActionReversalReason1Choice {
	c.Reason = new(CorporateActionReversalReason1Choice)
	return c.Reason
}

func (c *CorporateActionReversalReason1) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max256Text)(&value)
}
