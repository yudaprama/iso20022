package model

// Reason for the complete status.
type TransferCancellationCompleteReason1 struct {

	// Reason for a complete status in structured form.
	Structured *CancellationCompleteStatusReason1Code `xml:"Strd"`

	// Additional information about the reason for the complete status in textual form.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TransferCancellationCompleteReason1) SetStructured(value string) {
	t.Structured = (*CancellationCompleteStatusReason1Code)(&value)
}

func (t *TransferCancellationCompleteReason1) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}
