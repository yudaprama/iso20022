package model

// Choice between different quantity of security formats.
type Quantity21Choice struct {

	// Quantity of security.
	Quantity *FinancialInstrumentQuantity15Choice `xml:"Qty"`

	// Proprietary quantity of security format.
	ProprietaryQuantity *ProprietaryQuantity9 `xml:"PrtryQty"`
}

func (q *Quantity21Choice) AddQuantity() *FinancialInstrumentQuantity15Choice {
	q.Quantity = new(FinancialInstrumentQuantity15Choice)
	return q.Quantity
}

func (q *Quantity21Choice) AddProprietaryQuantity() *ProprietaryQuantity9 {
	q.ProprietaryQuantity = new(ProprietaryQuantity9)
	return q.ProprietaryQuantity
}
