package model

// Information about a rejected status.
type RejectionReason33 struct {

	// Reason for the rejected status.
	Reason *RejectedReason17Choice `xml:"Rsn"`

	// Additional information about the rejected status reason.
	AdditionalReasonInformation *Max350Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason33) AddReason() *RejectedReason17Choice {
	r.Reason = new(RejectedReason17Choice)
	return r.Reason
}

func (r *RejectionReason33) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max350Text)(&value)
}
