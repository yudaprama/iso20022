package model

// Specifies reasons for the rejected status.
type RejectedStatusReason22 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason23Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason22) AddReasonCode() *RejectedReason23Choice {
	r.ReasonCode = new(RejectedReason23Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason22) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
