package model

// Specifies the reason why the instruction or request has a rejected status.
//
type RejectionReason37 struct {

	// Reason provided for the status.
	Code *RejectionReason31Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason37) AddCode() *RejectionReason31Choice {
	r.Code = new(RejectionReason31Choice)
	return r.Code
}

func (r *RejectionReason37) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
