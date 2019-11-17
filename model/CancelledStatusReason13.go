package model

// Specifies reasons for the cancelled status.
type CancelledStatusReason13 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason10Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason13) AddReasonCode() *CancelledReason10Choice {
	c.ReasonCode = new(CancelledReason10Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason13) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
