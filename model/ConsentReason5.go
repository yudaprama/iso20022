package model

// Specifies the underlying reason for the status of an object.
type ConsentReason5 struct {

	// Reason provided for the status.
	Code *ConsentOrRejectionReason5Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *ConsentReason5) AddCode() *ConsentOrRejectionReason5Choice {
	c.Code = new(ConsentOrRejectionReason5Choice)
	return c.Code
}

func (c *ConsentReason5) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
