package model

// Reason for a rejected status.
type RejectedStatus10 struct {

	// Reason for the rejected status.
	Reason *RejectedReason21Choice `xml:"Rsn,omitempty"`

	// Additional information about the rejected reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectedStatus10) AddReason() *RejectedReason21Choice {
	r.Reason = new(RejectedReason21Choice)
	return r.Reason
}

func (r *RejectedStatus10) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max350Text)(&value)
}
