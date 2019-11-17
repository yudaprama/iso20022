package model

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason15 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason25Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason15) AddCode() *CancellationReason25Choice {
	c.Code = new(CancellationReason25Choice)
	return c.Code
}

func (c *CancellationReason15) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
