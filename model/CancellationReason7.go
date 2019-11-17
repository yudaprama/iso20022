package model

// The status of an instruction, advice or request.
type CancellationReason7 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason17Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason7) AddCode() *CancellationReason17Choice {
	c.Code = new(CancellationReason17Choice)
	return c.Code
}

func (c *CancellationReason7) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
