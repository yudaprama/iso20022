package model

// Specifies reasons for the cancelled status.
type CancelledStatusReason4 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason1Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason4) AddReasonCode() *CancelledReason1Choice {
	c.ReasonCode = new(CancelledReason1Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason4) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
