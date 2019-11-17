package model

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason25 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason23Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason25) AddCode() *RejectionReason23Choice {
	r.Code = new(RejectionReason23Choice)
	return r.Code
}

func (r *RejectionReason25) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
