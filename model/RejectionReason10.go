package model

// The status of an instruction, advice or request.
type RejectionReason10 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason10) AddCode() *RejectionReason10Choice {
	r.Code = new(RejectionReason10Choice)
	return r.Code
}

func (r *RejectionReason10) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
