package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason25 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason25Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason25) AddCode() *RejectionAndRepairReason25Choice {
	r.Code = new(RejectionAndRepairReason25Choice)
	return r.Code
}

func (r *RejectionOrRepairReason25) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
