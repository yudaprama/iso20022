package model

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason39 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason33Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason39) AddCode() *RejectionReason33Choice {
	r.Code = new(RejectionReason33Choice)
	return r.Code
}

func (r *RejectionReason39) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
