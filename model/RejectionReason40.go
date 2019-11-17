package model

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason40 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *ConsentOrRejectionReason5Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason40) AddCode() *ConsentOrRejectionReason5Choice {
	r.Code = new(ConsentOrRejectionReason5Choice)
	return r.Code
}

func (r *RejectionReason40) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
