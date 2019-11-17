package model

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason9 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason19Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason9) AddCode() *CancellationReason19Choice {
	c.Code = new(CancellationReason19Choice)
	return c.Code
}

func (c *CancellationReason9) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
