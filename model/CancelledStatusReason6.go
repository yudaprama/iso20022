package model

// Specifies reasons for the cancelled status.
type CancelledStatusReason6 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason3Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason6) AddReasonCode() *CancelledReason3Choice {
	c.ReasonCode = new(CancelledReason3Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason6) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
