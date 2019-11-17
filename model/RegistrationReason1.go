package model

// Reason of registration.
type RegistrationReason1 struct {

	// Specifies the reaoson of the holding status.
	Code *Registration3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalInformation *Max210Text `xml:"AddtlInf,omitempty"`
}

func (r *RegistrationReason1) AddCode() *Registration3Choice {
	r.Code = new(Registration3Choice)
	return r.Code
}

func (r *RegistrationReason1) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max210Text)(&value)
}
