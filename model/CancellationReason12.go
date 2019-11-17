package model

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason12 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason23Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason12) AddCode() *CancellationReason23Choice {
	c.Code = new(CancellationReason23Choice)
	return c.Code
}

func (c *CancellationReason12) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
