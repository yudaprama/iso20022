package model

// Provides reason of the rejection of a notification advice.
type NotificationRejectionReason1 struct {

	// The rejection reason.
	Reason []*RejectionReason6FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (n *NotificationRejectionReason1) AddReason() *RejectionReason6FormatChoice {
	newValue := new(RejectionReason6FormatChoice)
	n.Reason = append(n.Reason, newValue)
	return newValue
}

func (n *NotificationRejectionReason1) SetAdditionalInformation(value string) {
	n.AdditionalInformation = (*Max350Text)(&value)
}
