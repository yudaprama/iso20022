package model

// Transfer of a security on a securities account servicer after the receipt of a securities settlement instruction, or the transfer of cash on an account servicer after the receipt of a payment instruction. The transfer consists of a debit/credit operation in the books of the account servicer.
type TransferReference4 struct {

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`
}

func (t *TransferReference4) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferReference4) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}
