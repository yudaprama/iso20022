package model

// The status of an instruction, advice or request.
type CancellationReason1 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason5Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason1) AddCode() *CancellationReason5Choice {
	c.Code = new(CancellationReason5Choice)
	return c.Code
}

func (c *CancellationReason1) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
