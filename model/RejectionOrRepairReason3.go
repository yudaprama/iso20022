package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason3 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason3) AddCode() *RejectionAndRepairReason3Choice {
	r.Code = new(RejectionAndRepairReason3Choice)
	return r.Code
}

func (r *RejectionOrRepairReason3) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
