package model

// The status of an instruction, advice or request.
type RejectionReason18 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason15Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason18) AddCode() *RejectionReason15Choice {
	r.Code = new(RejectionReason15Choice)
	return r.Code
}

func (r *RejectionReason18) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
