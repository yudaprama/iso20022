package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason14 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code []*RejectionAndRepairReason14Choice `xml:"Cd,omitempty"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason14) AddCode() *RejectionAndRepairReason14Choice {
	newValue := new(RejectionAndRepairReason14Choice)
	r.Code = append(r.Code, newValue)
	return newValue
}

func (r *RejectionOrRepairReason14) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
