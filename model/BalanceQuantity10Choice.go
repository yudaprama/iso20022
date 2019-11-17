package model

// Choice between quantity formats for the balance information.
type BalanceQuantity10Choice struct {

	// Total quantity of financial instruments of the balance.
	Quantity *Quantity10Choice `xml:"Qty"`

	// Total quantity of financial instruments of the balance.
	Proprietary *GenericIdentification144 `xml:"Prtry"`
}

func (b *BalanceQuantity10Choice) AddQuantity() *Quantity10Choice {
	b.Quantity = new(Quantity10Choice)
	return b.Quantity
}

func (b *BalanceQuantity10Choice) AddProprietary() *GenericIdentification144 {
	b.Proprietary = new(GenericIdentification144)
	return b.Proprietary
}
