package model

// Specifies the underlying rejection reason for the status of an object.
type RejectionReason35 struct {

	// Specifies the reason why the instruction/request has a rejected status.
	Code *RejectionReason29Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason35) AddCode() *RejectionReason29Choice {
	r.Code = new(RejectionReason29Choice)
	return r.Code
}

func (r *RejectionReason35) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
