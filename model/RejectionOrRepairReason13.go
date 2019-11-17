package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason13 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason13Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason13) AddCode() *RejectionAndRepairReason13Choice {
	r.Code = new(RejectionAndRepairReason13Choice)
	return r.Code
}

func (r *RejectionOrRepairReason13) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
