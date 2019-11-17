package model

// Specifies reasons for the rejected status.
type RejectedStatusReason18 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason14Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason18) AddReasonCode() *RejectedReason14Choice {
	r.ReasonCode = new(RejectedReason14Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason18) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
