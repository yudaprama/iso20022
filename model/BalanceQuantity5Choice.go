package model

// Choice between quantity formats for the balance information.
type BalanceQuantity5Choice struct {

	// Total quantity of financial instruments of the balance.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Total quantity of financial instruments of the balance.
	Proprietary *GenericIdentification22 `xml:"Prtry"`
}

func (b *BalanceQuantity5Choice) AddQuantity() *FinancialInstrumentQuantity1Choice {
	b.Quantity = new(FinancialInstrumentQuantity1Choice)
	return b.Quantity
}

func (b *BalanceQuantity5Choice) AddProprietary() *GenericIdentification22 {
	b.Proprietary = new(GenericIdentification22)
	return b.Proprietary
}
