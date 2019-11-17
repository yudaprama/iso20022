package model

// Provides the status of a notification advice.
type NotificationAdviceStatus1Choice struct {

	// Provides information about the processing status of advice.
	ProcessedStatus *NotificationProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *NotificationRejectionReason1 `xml:"RjctdSts"`
}

func (n *NotificationAdviceStatus1Choice) AddProcessedStatus() *NotificationProcessingStatus1 {
	n.ProcessedStatus = new(NotificationProcessingStatus1)
	return n.ProcessedStatus
}

func (n *NotificationAdviceStatus1Choice) AddRejectedStatus() *NotificationRejectionReason1 {
	n.RejectedStatus = new(NotificationRejectionReason1)
	return n.RejectedStatus
}
