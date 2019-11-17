package model

// The status of an instruction, advice or request.
type CancellationReason5 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason12Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason5) AddCode() *CancellationReason12Choice {
	c.Code = new(CancellationReason12Choice)
	return c.Code
}

func (c *CancellationReason5) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
