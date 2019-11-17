package model

// Reason for an unmatched status.
type TransferUnmatchedStatusReason1 struct {

	// Reason for a rejected status in structured form.
	Structured *TransferUnmatchedReason1Code `xml:"Strd"`

	// Additional information about the reason for the unmatched status in textual form.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TransferUnmatchedStatusReason1) SetStructured(value string) {
	t.Structured = (*TransferUnmatchedReason1Code)(&value)
}

func (t *TransferUnmatchedStatusReason1) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}
