package model

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason28 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason26Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason28) AddCode() *RejectionReason26Choice {
	r.Code = new(RejectionReason26Choice)
	return r.Code
}

func (r *RejectionReason28) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
