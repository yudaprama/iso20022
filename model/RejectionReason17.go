package model

// Status of an instruction, advice or request.
type RejectionReason17 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason14Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason17) AddCode() *RejectionReason14Choice {
	r.Code = new(RejectionReason14Choice)
	return r.Code
}

func (r *RejectionReason17) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
