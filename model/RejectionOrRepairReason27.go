package model

// Reason for the rejection or repair status.
type RejectionOrRepairReason27 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code []*RejectionAndRepairReason27Choice `xml:"Cd,omitempty"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason27) AddCode() *RejectionAndRepairReason27Choice {
	newValue := new(RejectionAndRepairReason27Choice)
	r.Code = append(r.Code, newValue)
	return newValue
}

func (r *RejectionOrRepairReason27) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
