package model

// Reason of registration.
type RegistrationReason6 struct {

	// Specifies the reaoson of the holding status.
	Code *Registration12Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalInformation *RestrictedFINXMax210Text `xml:"AddtlInf,omitempty"`
}

func (r *RegistrationReason6) AddCode() *Registration12Choice {
	r.Code = new(Registration12Choice)
	return r.Code
}

func (r *RegistrationReason6) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*RestrictedFINXMax210Text)(&value)
}
