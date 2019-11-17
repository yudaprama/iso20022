package model

// Information about a transfer in transaction.
type TransferIn11 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer21 `xml:"TrfDtls"`
}

func (t *TransferIn11) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferIn11) AddTransferDetails() *Transfer21 {
	newValue := new(Transfer21)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}
