package model

// The status of an instruction, advice or request.
type RejectionReason9 struct {

	// Specifies the reason why the instruction/request has a rejected status.
	Code *RejectionReason9Choice `xml:"Cd"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason9) AddCode() *RejectionReason9Choice {
	r.Code = new(RejectionReason9Choice)
	return r.Code
}

func (r *RejectionReason9) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
