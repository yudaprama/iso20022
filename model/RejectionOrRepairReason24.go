package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason24 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason24Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason24) AddCode() *RejectionAndRepairReason24Choice {
	r.Code = new(RejectionAndRepairReason24Choice)
	return r.Code
}

func (r *RejectionOrRepairReason24) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
