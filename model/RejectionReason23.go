package model

// Reason to reject the message.
type RejectionReason23 struct {

	// Reason to reject the message.
	Reason *MessageRejectedReason1Code `xml:"Rsn"`

	// Additional information about the rejection reason.
	AdditionalInformation *Max140Text `xml:"AddtlInf,omitempty"`

	// Identification of the invalid or unrecognised reference.
	LinkedMessage *LinkedMessage1Choice `xml:"LkdMsg,omitempty"`
}

func (r *RejectionReason23) SetReason(value string) {
	r.Reason = (*MessageRejectedReason1Code)(&value)
}

func (r *RejectionReason23) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max140Text)(&value)
}

func (r *RejectionReason23) AddLinkedMessage() *LinkedMessage1Choice {
	r.LinkedMessage = new(LinkedMessage1Choice)
	return r.LinkedMessage
}
