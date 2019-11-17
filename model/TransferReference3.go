package model

// Reference of a transfer instruction cancellation.
type TransferReference3 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Transfer and cancellation reference.
	TransferReferences []*TransferReference4 `xml:"TrfRefs"`
}

func (t *TransferReference3) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferReference3) AddTransferReferences() *TransferReference4 {
	newValue := new(TransferReference4)
	t.TransferReferences = append(t.TransferReferences, newValue)
	return newValue
}
