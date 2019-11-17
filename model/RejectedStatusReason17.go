package model

// Specifies reasons for the rejected status.
type RejectedStatusReason17 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason13Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason17) AddReasonCode() *RejectedReason13Choice {
	r.ReasonCode = new(RejectedReason13Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason17) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
