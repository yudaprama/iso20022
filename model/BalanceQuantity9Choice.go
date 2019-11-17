package model

// Choice between quantity formats for the balance information.
type BalanceQuantity9Choice struct {

	// Total quantity of financial instruments of the balance.
	Quantity *Quantity6Choice `xml:"Qty"`

	// Total quantity of financial instruments of the balance.
	Proprietary *GenericIdentification56 `xml:"Prtry"`
}

func (b *BalanceQuantity9Choice) AddQuantity() *Quantity6Choice {
	b.Quantity = new(Quantity6Choice)
	return b.Quantity
}

func (b *BalanceQuantity9Choice) AddProprietary() *GenericIdentification56 {
	b.Proprietary = new(GenericIdentification56)
	return b.Proprietary
}
