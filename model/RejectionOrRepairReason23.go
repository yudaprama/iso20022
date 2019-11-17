package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason23 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code []*RejectionAndRepairReason23Choice `xml:"Cd,omitempty"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason23) AddCode() *RejectionAndRepairReason23Choice {
	newValue := new(RejectionAndRepairReason23Choice)
	r.Code = append(r.Code, newValue)
	return newValue
}

func (r *RejectionOrRepairReason23) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
