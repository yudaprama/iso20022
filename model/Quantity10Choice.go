package model

// Choice of format for the quantity.
type Quantity10Choice struct {

	// Quantity of financial instrument in units, original face amount or current face amount.
	Quantity *FinancialInstrumentQuantity15Choice `xml:"Qty"`

	// Original and current value of an asset-back instrument.
	OriginalAndCurrentFace *OriginalAndCurrentQuantities4 `xml:"OrgnlAndCurFace"`
}

func (q *Quantity10Choice) AddQuantity() *FinancialInstrumentQuantity15Choice {
	q.Quantity = new(FinancialInstrumentQuantity15Choice)
	return q.Quantity
}

func (q *Quantity10Choice) AddOriginalAndCurrentFace() *OriginalAndCurrentQuantities4 {
	q.OriginalAndCurrentFace = new(OriginalAndCurrentQuantities4)
	return q.OriginalAndCurrentFace
}
