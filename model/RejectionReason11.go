package model

// The status of an instruction, advice or request.
type RejectionReason11 struct {

	// Reason provided for the status.
	Code *RejectionReason11Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason11) AddCode() *RejectionReason11Choice {
	r.Code = new(RejectionReason11Choice)
	return r.Code
}

func (r *RejectionReason11) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
