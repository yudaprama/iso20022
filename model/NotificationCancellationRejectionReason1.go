package model

// Provides reason of the rejection of a notification cancellation request.
type NotificationCancellationRejectionReason1 struct {

	// The rejection reason.
	Reason []*RejectionReason11FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (n *NotificationCancellationRejectionReason1) AddReason() *RejectionReason11FormatChoice {
	newValue := new(RejectionReason11FormatChoice)
	n.Reason = append(n.Reason, newValue)
	return newValue
}

func (n *NotificationCancellationRejectionReason1) SetAdditionalInformation(value string) {
	n.AdditionalInformation = (*Max350Text)(&value)
}
