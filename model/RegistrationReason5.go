package model

// Reason of registration.
type RegistrationReason5 struct {

	// Specifies the reaoson of the holding status.
	Code *Registration10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalInformation *Max210Text `xml:"AddtlInf,omitempty"`
}

func (r *RegistrationReason5) AddCode() *Registration10Choice {
	r.Code = new(Registration10Choice)
	return r.Code
}

func (r *RegistrationReason5) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max210Text)(&value)
}
