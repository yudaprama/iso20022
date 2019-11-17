package model

// Reason of registration.
type RegistrationReason3 struct {

	// Specifies the reaoson of the holding status.
	Code *Registration7Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalInformation *Max210Text `xml:"AddtlInf,omitempty"`
}

func (r *RegistrationReason3) AddCode() *Registration7Choice {
	r.Code = new(Registration7Choice)
	return r.Code
}

func (r *RegistrationReason3) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max210Text)(&value)
}
