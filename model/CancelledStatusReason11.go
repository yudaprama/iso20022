package model

// Specifies reasons for the cancelled status.
type CancelledStatusReason11 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason8Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason11) AddReasonCode() *CancelledReason8Choice {
	c.ReasonCode = new(CancelledReason8Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason11) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
