package model

// Choice between quantity formats for the balance information.
type BalanceQuantity12Choice struct {

	// Total quantity of financial instruments of the balance.
	Quantity *FinancialInstrumentQuantity15Choice `xml:"Qty"`

	// Total quantity of financial instruments of the balance.
	Proprietary *GenericIdentification144 `xml:"Prtry"`
}

func (b *BalanceQuantity12Choice) AddQuantity() *FinancialInstrumentQuantity15Choice {
	b.Quantity = new(FinancialInstrumentQuantity15Choice)
	return b.Quantity
}

func (b *BalanceQuantity12Choice) AddProprietary() *GenericIdentification144 {
	b.Proprietary = new(GenericIdentification144)
	return b.Proprietary
}
