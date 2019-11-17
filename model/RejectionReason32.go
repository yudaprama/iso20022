package model

// Information about a rejected status.
type RejectionReason32 struct {

	// Reason for the rejected status.
	Reason *RejectedReason15Choice `xml:"Rsn"`

	// Additional information about the rejected status reason.
	AdditionalReasonInformation *Max350Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason32) AddReason() *RejectedReason15Choice {
	r.Reason = new(RejectedReason15Choice)
	return r.Reason
}

func (r *RejectionReason32) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max350Text)(&value)
}
