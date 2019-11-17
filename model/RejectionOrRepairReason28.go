package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason28 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason28Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason28) AddCode() *RejectionAndRepairReason28Choice {
	r.Code = new(RejectionAndRepairReason28Choice)
	return r.Code
}

func (r *RejectionOrRepairReason28) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
