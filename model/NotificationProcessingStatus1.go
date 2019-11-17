package model

// Provide processing status information for a notification advice.
type NotificationProcessingStatus1 struct {

	// The processing status.
	Status *ProcessedStatus1FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (n *NotificationProcessingStatus1) AddStatus() *ProcessedStatus1FormatChoice {
	n.Status = new(ProcessedStatus1FormatChoice)
	return n.Status
}

func (n *NotificationProcessingStatus1) SetAdditionalInformation(value string) {
	n.AdditionalInformation = (*Max350Text)(&value)
}
