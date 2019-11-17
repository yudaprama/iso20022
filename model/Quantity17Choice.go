package model

// Choice between different quantity of security formats.
type Quantity17Choice struct {

	// Choice between different quantity of security formats.
	QuantityChoice *Quantity18Choice `xml:"QtyChc"`

	// Proprietary quantity of security format.
	ProprietaryQuantity *ProprietaryQuantity7 `xml:"PrtryQty"`
}

func (q *Quantity17Choice) AddQuantityChoice() *Quantity18Choice {
	q.QuantityChoice = new(Quantity18Choice)
	return q.QuantityChoice
}

func (q *Quantity17Choice) AddProprietaryQuantity() *ProprietaryQuantity7 {
	q.ProprietaryQuantity = new(ProprietaryQuantity7)
	return q.ProprietaryQuantity
}
