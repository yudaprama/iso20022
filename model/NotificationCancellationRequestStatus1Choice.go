package model

// Choice between various statuses.
type NotificationCancellationRequestStatus1Choice struct {

	// Provides information about the processing status of the cancellation request.
	ProcessedStatus *NotificationCancellationProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *NotificationCancellationRejectionReason1 `xml:"RjctdSts"`
}

func (n *NotificationCancellationRequestStatus1Choice) AddProcessedStatus() *NotificationCancellationProcessingStatus1 {
	n.ProcessedStatus = new(NotificationCancellationProcessingStatus1)
	return n.ProcessedStatus
}

func (n *NotificationCancellationRequestStatus1Choice) AddRejectedStatus() *NotificationCancellationRejectionReason1 {
	n.RejectedStatus = new(NotificationCancellationRejectionReason1)
	return n.RejectedStatus
}
