package model

// Reference of a transfer instruction cancellation.
type TransferReference7 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Transfer and cancellation reference.
	TransferReferences []*TransferReference8 `xml:"TrfRefs"`
}

func (t *TransferReference7) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferReference7) AddTransferReferences() *TransferReference8 {
	newValue := new(TransferReference8)
	t.TransferReferences = append(t.TransferReferences, newValue)
	return newValue
}
