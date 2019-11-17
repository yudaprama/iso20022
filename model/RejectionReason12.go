package model

// Specifies the underlying reason for the status of an object.
type RejectionReason12 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *ConsentOrRejectionReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason12) AddCode() *ConsentOrRejectionReason2Choice {
	r.Code = new(ConsentOrRejectionReason2Choice)
	return r.Code
}

func (r *RejectionReason12) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
