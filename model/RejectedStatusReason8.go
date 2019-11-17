package model

// Specifies reasons for the rejected status.
type RejectedStatusReason8 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason1Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason8) AddReasonCode() *RejectedReason1Choice {
	r.ReasonCode = new(RejectedReason1Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason8) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
