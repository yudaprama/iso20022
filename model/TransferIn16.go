package model

// Information about a transfer in transaction.
type TransferIn16 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer32 `xml:"TrfDtls"`
}

func (t *TransferIn16) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferIn16) AddTransferDetails() *Transfer32 {
	newValue := new(Transfer32)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}
