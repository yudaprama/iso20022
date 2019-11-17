package model

// Choice between different quantity of security formats.
type Quantity2Choice struct {

	// Quantity of security.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Proprietary quantity of security format.
	ProprietaryQuantity *ProprietaryQuantity2 `xml:"PrtryQty"`
}

func (q *Quantity2Choice) AddQuantity() *FinancialInstrumentQuantity1Choice {
	q.Quantity = new(FinancialInstrumentQuantity1Choice)
	return q.Quantity
}

func (q *Quantity2Choice) AddProprietaryQuantity() *ProprietaryQuantity2 {
	q.ProprietaryQuantity = new(ProprietaryQuantity2)
	return q.ProprietaryQuantity
}
