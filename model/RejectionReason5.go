package model

// The status of an instruction, advice or request.
type RejectionReason5 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason4Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason5) AddCode() *RejectionReason4Choice {
	r.Code = new(RejectionReason4Choice)
	return r.Code
}

func (r *RejectionReason5) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
