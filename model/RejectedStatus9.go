package model

// Reason for a rejected status.
type RejectedStatus9 struct {

	// Reason for the rejected status.
	Reason *RejectedReason20Choice `xml:"Rsn,omitempty"`

	// Additional information about the rejected reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectedStatus9) AddReason() *RejectedReason20Choice {
	r.Reason = new(RejectedReason20Choice)
	return r.Reason
}

func (r *RejectedStatus9) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max350Text)(&value)
}
