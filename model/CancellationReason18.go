package model

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason18 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason28Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason18) AddCode() *CancellationReason28Choice {
	c.Code = new(CancellationReason28Choice)
	return c.Code
}

func (c *CancellationReason18) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
