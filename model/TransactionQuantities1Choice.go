package model

// Specifies the quantities (eg of securities) in the underlying transaction.
type TransactionQuantities1Choice struct {

	// Specifies the quantity (eg of securities) in the underlying transaction.
	Quantity *FinancialInstrumentQuantityChoice `xml:"Qty"`

	// Proprietary quantities specification defined in the underlying transaction.
	Proprietary *ProprietaryQuantity1 `xml:"Prtry"`
}

func (t *TransactionQuantities1Choice) AddQuantity() *FinancialInstrumentQuantityChoice {
	t.Quantity = new(FinancialInstrumentQuantityChoice)
	return t.Quantity
}

func (t *TransactionQuantities1Choice) AddProprietary() *ProprietaryQuantity1 {
	t.Proprietary = new(ProprietaryQuantity1)
	return t.Proprietary
}
