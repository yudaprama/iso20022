package model

// Choice between different formats of securities quantity.
type Quantity20Choice struct {

	// Standard code to specify quantity of a financial instrument.
	Code *Quantity1Code `xml:"Cd"`

	// Face amount and amortised value of security.
	OriginalAndCurrentFaceAmount *OriginalAndCurrentQuantities1 `xml:"OrgnlAndCurFaceAmt"`

	// Quantity of financial instrument in units, original face amount or current face amount.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`
}

func (q *Quantity20Choice) SetCode(value string) {
	q.Code = (*Quantity1Code)(&value)
}

func (q *Quantity20Choice) AddOriginalAndCurrentFaceAmount() *OriginalAndCurrentQuantities1 {
	q.OriginalAndCurrentFaceAmount = new(OriginalAndCurrentQuantities1)
	return q.OriginalAndCurrentFaceAmount
}

func (q *Quantity20Choice) AddQuantity() *FinancialInstrumentQuantity1Choice {
	q.Quantity = new(FinancialInstrumentQuantity1Choice)
	return q.Quantity
}
