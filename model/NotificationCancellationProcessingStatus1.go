package model

// Provide processing status information for a notification cancellation request.
type NotificationCancellationProcessingStatus1 struct {

	// The processing status.
	Status *ProcessedStatus2FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (n *NotificationCancellationProcessingStatus1) AddStatus() *ProcessedStatus2FormatChoice {
	n.Status = new(ProcessedStatus2FormatChoice)
	return n.Status
}

func (n *NotificationCancellationProcessingStatus1) SetAdditionalInformation(value string) {
	n.AdditionalInformation = (*Max350Text)(&value)
}
