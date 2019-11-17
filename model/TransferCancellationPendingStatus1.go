package model

// Transfer cancellation status is in pending.
type TransferCancellationPendingStatus1 struct {

	// Reason for the cancellation pending status.
	Reason *Max350Text `xml:"Rsn,omitempty"`
}

func (t *TransferCancellationPendingStatus1) SetReason(value string) {
	t.Reason = (*Max350Text)(&value)
}
