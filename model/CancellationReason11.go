package model

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason11 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason22Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason11) AddCode() *CancellationReason22Choice {
	c.Code = new(CancellationReason22Choice)
	return c.Code
}

func (c *CancellationReason11) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
