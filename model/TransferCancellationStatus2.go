package model

// Transfer cancellation status is accepted or sent to next party.
type TransferCancellationStatus2 struct {

	// Status of the transfer cancellation is accepted or sent to next party.
	Status *CancellationStatus2Code `xml:"Sts"`

	// Reason for the status.
	Reason *Max350Text `xml:"Rsn,omitempty"`
}

func (t *TransferCancellationStatus2) SetStatus(value string) {
	t.Status = (*CancellationStatus2Code)(&value)
}

func (t *TransferCancellationStatus2) SetReason(value string) {
	t.Reason = (*Max350Text)(&value)
}
