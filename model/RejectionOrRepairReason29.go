package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason29 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason29Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason29) AddCode() *RejectionAndRepairReason29Choice {
	r.Code = new(RejectionAndRepairReason29Choice)
	return r.Code
}

func (r *RejectionOrRepairReason29) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
