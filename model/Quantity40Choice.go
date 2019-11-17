package model

// Choice between different quantity of security formats.
type Quantity40Choice struct {

	// Standard code to specify quantity of a financial instrument.
	Code *Quantity1Code `xml:"Cd"`

	// Face amount and amortised value of security.
	OriginalAndCurrentFaceAmount *OriginalAndCurrentQuantities4 `xml:"OrgnlAndCurFaceAmt"`

	// Quantity of financial instrument in units, original face amount or current face amount.
	Quantity *FinancialInstrumentQuantity15Choice `xml:"Qty"`
}

func (q *Quantity40Choice) SetCode(value string) {
	q.Code = (*Quantity1Code)(&value)
}

func (q *Quantity40Choice) AddOriginalAndCurrentFaceAmount() *OriginalAndCurrentQuantities4 {
	q.OriginalAndCurrentFaceAmount = new(OriginalAndCurrentQuantities4)
	return q.OriginalAndCurrentFaceAmount
}

func (q *Quantity40Choice) AddQuantity() *FinancialInstrumentQuantity15Choice {
	q.Quantity = new(FinancialInstrumentQuantity15Choice)
	return q.Quantity
}
