package model

// Information about a rejected status.
type RejectionReason31 struct {

	// Reason for the rejected status.
	Reason *RejectedReason16Choice `xml:"Rsn"`

	// Additional information about the rejected status reason.
	AdditionalReasonInformation *Max350Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason31) AddReason() *RejectedReason16Choice {
	r.Reason = new(RejectedReason16Choice)
	return r.Reason
}

func (r *RejectionReason31) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max350Text)(&value)
}
