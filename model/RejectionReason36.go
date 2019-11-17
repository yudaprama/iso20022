package model

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason36 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason30Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason36) AddCode() *RejectionReason30Choice {
	r.Code = new(RejectionReason30Choice)
	return r.Code
}

func (r *RejectionReason36) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
