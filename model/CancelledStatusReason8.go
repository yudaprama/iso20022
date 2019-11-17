package model

// Specifies reasons for the cancelled status.
type CancelledStatusReason8 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason5Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason8) AddReasonCode() *CancelledReason5Choice {
	c.ReasonCode = new(CancelledReason5Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason8) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
