package model

// Describes the type of product and the assets to be transferred.
type ISATransfer20 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*ISATransfer18 `xml:"PdctTrf"`
}

func (i *ISATransfer20) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *ISATransfer20) AddProductTransfer() *ISATransfer18 {
	newValue := new(ISATransfer18)
	i.ProductTransfer = append(i.ProductTransfer, newValue)
	return newValue
}
