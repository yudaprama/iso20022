package model

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason14 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason24Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason14) AddCode() *CancellationReason24Choice {
	c.Code = new(CancellationReason24Choice)
	return c.Code
}

func (c *CancellationReason14) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
