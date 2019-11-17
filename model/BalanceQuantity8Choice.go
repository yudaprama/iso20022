package model

// Choice between quantity formats for the balance information.
type BalanceQuantity8Choice struct {

	// Total quantity of financial instruments of the balance.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Total quantity of financial instruments of the balance.
	Proprietary *GenericIdentification56 `xml:"Prtry"`
}

func (b *BalanceQuantity8Choice) AddQuantity() *FinancialInstrumentQuantity1Choice {
	b.Quantity = new(FinancialInstrumentQuantity1Choice)
	return b.Quantity
}

func (b *BalanceQuantity8Choice) AddProprietary() *GenericIdentification56 {
	b.Proprietary = new(GenericIdentification56)
	return b.Proprietary
}
