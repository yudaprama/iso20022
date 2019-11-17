package model

// Reason for the rejected status.
type TransferCancellationRejectionReason1 struct {

	// Reason for a rejected status in structured form.
	Structured *CancellationRejectedReason1Code `xml:"Strd"`

	// Additional information about the reason for the rejected status in textual form.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TransferCancellationRejectionReason1) SetStructured(value string) {
	t.Structured = (*CancellationRejectedReason1Code)(&value)
}

func (t *TransferCancellationRejectionReason1) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}
