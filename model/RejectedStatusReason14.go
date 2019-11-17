package model

// Specifies reasons for the rejected status.
type RejectedStatusReason14 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason10Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason14) AddReasonCode() *RejectedReason10Choice {
	r.ReasonCode = new(RejectedReason10Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason14) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
