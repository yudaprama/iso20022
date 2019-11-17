package model

// Information about a transfer out transaction.
type TransferOut13 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer27 `xml:"TrfDtls"`
}

func (t *TransferOut13) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferOut13) AddTransferDetails() *Transfer27 {
	newValue := new(Transfer27)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}
