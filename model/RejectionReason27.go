package model

// Specifies the underlying rejection reason for the status of an object.
type RejectionReason27 struct {

	// Specifies the reason why the instruction/request has a rejected status.
	Code *RejectionReason25Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason27) AddCode() *RejectionReason25Choice {
	r.Code = new(RejectionReason25Choice)
	return r.Code
}

func (r *RejectionReason27) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
