package model

// Choice between different quantity of security formats.
type Quantity19Choice struct {

	// Quantity of security.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Proprietary quantity of security format.
	ProprietaryQuantity *ProprietaryQuantity8 `xml:"PrtryQty"`
}

func (q *Quantity19Choice) AddQuantity() *FinancialInstrumentQuantity1Choice {
	q.Quantity = new(FinancialInstrumentQuantity1Choice)
	return q.Quantity
}

func (q *Quantity19Choice) AddProprietaryQuantity() *ProprietaryQuantity8 {
	q.ProprietaryQuantity = new(ProprietaryQuantity8)
	return q.ProprietaryQuantity
}
