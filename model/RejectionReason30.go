package model

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason30 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason27Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason30) AddCode() *RejectionReason27Choice {
	r.Code = new(RejectionReason27Choice)
	return r.Code
}

func (r *RejectionReason30) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
