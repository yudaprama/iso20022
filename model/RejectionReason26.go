package model

// Specifies the reason why the instruction or request has a rejected status.
//
type RejectionReason26 struct {

	// Reason provided for the status.
	Code *RejectionReason24Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason26) AddCode() *RejectionReason24Choice {
	r.Code = new(RejectionReason24Choice)
	return r.Code
}

func (r *RejectionReason26) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
