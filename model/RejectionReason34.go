package model

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason34 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason28Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason34) AddCode() *RejectionReason28Choice {
	r.Code = new(RejectionReason28Choice)
	return r.Code
}

func (r *RejectionReason34) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
