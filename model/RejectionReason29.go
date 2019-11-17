package model

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason29 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *ConsentOrRejectionReason4Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason29) AddCode() *ConsentOrRejectionReason4Choice {
	r.Code = new(ConsentOrRejectionReason4Choice)
	return r.Code
}

func (r *RejectionReason29) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
