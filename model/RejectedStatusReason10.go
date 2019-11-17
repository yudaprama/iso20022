package model

// Specifies reasons for the rejected status.
type RejectedStatusReason10 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason5Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason10) AddReasonCode() *RejectedReason5Choice {
	r.ReasonCode = new(RejectedReason5Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason10) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
