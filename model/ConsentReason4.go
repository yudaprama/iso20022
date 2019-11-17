package model

// Specifies the underlying reason for the status of an object.
type ConsentReason4 struct {

	// Reason provided for the status.
	Code *ConsentOrRejectionReason4Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *ConsentReason4) AddCode() *ConsentOrRejectionReason4Choice {
	c.Code = new(ConsentOrRejectionReason4Choice)
	return c.Code
}

func (c *ConsentReason4) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
