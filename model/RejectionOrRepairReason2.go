package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason2 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason2) AddCode() *RejectionAndRepairReason2Choice {
	r.Code = new(RejectionAndRepairReason2Choice)
	return r.Code
}

func (r *RejectionOrRepairReason2) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
