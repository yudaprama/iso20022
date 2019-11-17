package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason31 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason31Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason31) AddCode() *RejectionAndRepairReason31Choice {
	r.Code = new(RejectionAndRepairReason31Choice)
	return r.Code
}

func (r *RejectionOrRepairReason31) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
