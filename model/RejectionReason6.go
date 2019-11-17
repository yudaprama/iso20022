package model

// The status of an instruction, advice or request.
type RejectionReason6 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason6) AddCode() *RejectionReason2Choice {
	r.Code = new(RejectionReason2Choice)
	return r.Code
}

func (r *RejectionReason6) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
