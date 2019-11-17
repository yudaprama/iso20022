package model

// Choice between formats for the balance information.
type BalanceQuantity1Choice struct {

	// Total quantity of financial instruments of the balance.
	Quantity *FinancialInstrumentQuantityChoice `xml:"Qty"`

	// Total quantity of financial instruments of the balance.
	QuantityAsDSS *GenericIdentification6 `xml:"QtyAsDSS"`
}

func (b *BalanceQuantity1Choice) AddQuantity() *FinancialInstrumentQuantityChoice {
	b.Quantity = new(FinancialInstrumentQuantityChoice)
	return b.Quantity
}

func (b *BalanceQuantity1Choice) AddQuantityAsDSS() *GenericIdentification6 {
	b.QuantityAsDSS = new(GenericIdentification6)
	return b.QuantityAsDSS
}
