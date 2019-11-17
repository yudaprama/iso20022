package model

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason10 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason21Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason10) AddCode() *CancellationReason21Choice {
	c.Code = new(CancellationReason21Choice)
	return c.Code
}

func (c *CancellationReason10) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
