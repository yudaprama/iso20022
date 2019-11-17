package model

// Describes the type of product and the assets to be transferred.
type ISATransfer17 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*ISATransfer16 `xml:"PdctTrf"`
}

func (i *ISATransfer17) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *ISATransfer17) AddProductTransfer() *ISATransfer16 {
	newValue := new(ISATransfer16)
	i.ProductTransfer = append(i.ProductTransfer, newValue)
	return newValue
}
