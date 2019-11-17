package model

// The status of an instruction, advice or request.
type CancellationReason2 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason2) AddCode() *CancellationReason3Choice {
	c.Code = new(CancellationReason3Choice)
	return c.Code
}

func (c *CancellationReason2) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
