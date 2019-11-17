package model

// Specifies the underlying rejection reason for the status of an object.
type RejectionReason19 struct {

	// Specifies the reason why the instruction/request has a rejected status.
	Code *RejectionReason17Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason19) AddCode() *RejectionReason17Choice {
	r.Code = new(RejectionReason17Choice)
	return r.Code
}

func (r *RejectionReason19) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
