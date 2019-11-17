package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason1 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason1) AddCode() *RejectionAndRepairReason1Choice {
	r.Code = new(RejectionAndRepairReason1Choice)
	return r.Code
}

func (r *RejectionOrRepairReason1) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
