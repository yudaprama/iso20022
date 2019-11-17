package model

// Choice of format for the quantity.
type Quantity6Choice struct {

	// Quantity of financial instrument in units, original face amount or current face amount.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Original and current value of an asset-back instrument.
	OriginalAndCurrentFace *OriginalAndCurrentQuantities1 `xml:"OrgnlAndCurFace"`
}

func (q *Quantity6Choice) AddQuantity() *FinancialInstrumentQuantity1Choice {
	q.Quantity = new(FinancialInstrumentQuantity1Choice)
	return q.Quantity
}

func (q *Quantity6Choice) AddOriginalAndCurrentFace() *OriginalAndCurrentQuantities1 {
	q.OriginalAndCurrentFace = new(OriginalAndCurrentQuantities1)
	return q.OriginalAndCurrentFace
}
