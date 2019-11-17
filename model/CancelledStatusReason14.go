package model

// Specifies reasons for the cancelled status.
type CancelledStatusReason14 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason11Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason14) AddReasonCode() *CancelledReason11Choice {
	c.ReasonCode = new(CancelledReason11Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason14) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
