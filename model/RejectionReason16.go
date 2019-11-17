package model

// Information about a rejected status.
type RejectionReason16 struct {

	// Reason for the rejected status.
	Reason *RejectedReason4Choice `xml:"Rsn"`

	// Additional information about the rejected status reason.
	AdditionalReasonInformation *Max350Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason16) AddReason() *RejectedReason4Choice {
	r.Reason = new(RejectedReason4Choice)
	return r.Reason
}

func (r *RejectionReason16) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max350Text)(&value)
}
