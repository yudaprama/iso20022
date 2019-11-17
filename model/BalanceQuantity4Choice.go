package model

// Choice between quantity formats for the balance information.
type BalanceQuantity4Choice struct {

	// Total quantity of financial instruments of the balance.
	Quantity *Quantity6Choice `xml:"Qty"`

	// Total quantity of financial instruments of the balance.
	Proprietary *GenericIdentification22 `xml:"Prtry"`
}

func (b *BalanceQuantity4Choice) AddQuantity() *Quantity6Choice {
	b.Quantity = new(Quantity6Choice)
	return b.Quantity
}

func (b *BalanceQuantity4Choice) AddProprietary() *GenericIdentification22 {
	b.Proprietary = new(GenericIdentification22)
	return b.Proprietary
}
