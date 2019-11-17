package model

// Specifies the underlying reason for the status of an object.
type ConsentReason2 struct {

	// Reason provided for the status.
	Code *ConsentOrRejectionReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *ConsentReason2) AddCode() *ConsentOrRejectionReason2Choice {
	c.Code = new(ConsentOrRejectionReason2Choice)
	return c.Code
}

func (c *ConsentReason2) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
