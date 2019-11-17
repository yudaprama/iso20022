package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason18 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason18Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason18) AddCode() *RejectionAndRepairReason18Choice {
	r.Code = new(RejectionAndRepairReason18Choice)
	return r.Code
}

func (r *RejectionOrRepairReason18) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
