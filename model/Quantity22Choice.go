package model

// Choice between different quantity of security formats.
type Quantity22Choice struct {

	// Choice between different quantity of security formats.
	QuantityChoice *Quantity23Choice `xml:"QtyChc"`

	// Proprietary quantity of security format.
	ProprietaryQuantity *ProprietaryQuantity10 `xml:"PrtryQty"`
}

func (q *Quantity22Choice) AddQuantityChoice() *Quantity23Choice {
	q.QuantityChoice = new(Quantity23Choice)
	return q.QuantityChoice
}

func (q *Quantity22Choice) AddProprietaryQuantity() *ProprietaryQuantity10 {
	q.ProprietaryQuantity = new(ProprietaryQuantity10)
	return q.ProprietaryQuantity
}
