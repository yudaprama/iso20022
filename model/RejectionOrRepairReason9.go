package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason9 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason9Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason9) AddCode() *RejectionAndRepairReason9Choice {
	r.Code = new(RejectionAndRepairReason9Choice)
	return r.Code
}

func (r *RejectionOrRepairReason9) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
