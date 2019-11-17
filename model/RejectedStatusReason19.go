package model

// Specifies reasons for the rejected status.
type RejectedStatusReason19 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason18Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason19) AddReasonCode() *RejectedReason18Choice {
	r.ReasonCode = new(RejectedReason18Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason19) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
