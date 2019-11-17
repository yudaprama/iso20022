package model

// Choice between different quantity of security formats.
type Quantity3Choice struct {

	// Choice between different quantity of security formats.
	QuantityChoice *Quantity4Choice `xml:"QtyChc"`

	// Proprietary quantity of security format.
	ProprietaryQuantity *ProprietaryQuantity3 `xml:"PrtryQty"`
}

func (q *Quantity3Choice) AddQuantityChoice() *Quantity4Choice {
	q.QuantityChoice = new(Quantity4Choice)
	return q.QuantityChoice
}

func (q *Quantity3Choice) AddProprietaryQuantity() *ProprietaryQuantity3 {
	q.ProprietaryQuantity = new(ProprietaryQuantity3)
	return q.ProprietaryQuantity
}
