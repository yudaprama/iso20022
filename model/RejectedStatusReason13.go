package model

// Specifies reasons for the rejected status.
type RejectedStatusReason13 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason9Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason13) AddReasonCode() *RejectedReason9Choice {
	r.ReasonCode = new(RejectedReason9Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason13) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
