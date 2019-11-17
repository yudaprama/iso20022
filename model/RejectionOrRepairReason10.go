package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason10 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason10) AddCode() *RejectionAndRepairReason10Choice {
	r.Code = new(RejectionAndRepairReason10Choice)
	return r.Code
}

func (r *RejectionOrRepairReason10) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
