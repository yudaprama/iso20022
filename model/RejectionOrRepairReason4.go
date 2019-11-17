package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason4 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code []*RejectionAndRepairReason4Choice `xml:"Cd,omitempty"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason4) AddCode() *RejectionAndRepairReason4Choice {
	newValue := new(RejectionAndRepairReason4Choice)
	r.Code = append(r.Code, newValue)
	return newValue
}

func (r *RejectionOrRepairReason4) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
