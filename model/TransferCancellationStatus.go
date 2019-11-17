package model

// Cancellation status.
type TransferCancellationStatus struct {

	// Status of the transfer cancellation instruction.
	Status *CancellationStatus1Code `xml:"Sts"`

	// Additional information about the status in textual form.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TransferCancellationStatus) SetStatus(value string) {
	t.Status = (*CancellationStatus1Code)(&value)
}

func (t *TransferCancellationStatus) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}
