package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason26 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason26Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason26) AddCode() *RejectionAndRepairReason26Choice {
	r.Code = new(RejectionAndRepairReason26Choice)
	return r.Code
}

func (r *RejectionOrRepairReason26) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
