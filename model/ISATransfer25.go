package model

// Describes the type of product and the assets to be transferred.
type ISATransfer25 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*ISATransfer22 `xml:"PdctTrf"`
}

func (i *ISATransfer25) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *ISATransfer25) AddProductTransfer() *ISATransfer22 {
	newValue := new(ISATransfer22)
	i.ProductTransfer = append(i.ProductTransfer, newValue)
	return newValue
}
