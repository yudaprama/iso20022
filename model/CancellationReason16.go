package model

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason16 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason26Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason16) AddCode() *CancellationReason26Choice {
	c.Code = new(CancellationReason26Choice)
	return c.Code
}

func (c *CancellationReason16) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
