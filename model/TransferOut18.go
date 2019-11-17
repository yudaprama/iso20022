package model

// Information about a transfer out transaction.
type TransferOut18 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer30 `xml:"TrfDtls"`
}

func (t *TransferOut18) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferOut18) AddTransferDetails() *Transfer30 {
	newValue := new(Transfer30)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}
