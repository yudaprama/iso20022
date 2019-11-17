package model

// Specifies the quantities (such as securities) in the underlying transaction.
type TransactionQuantities2Choice struct {

	// Specifies the quantity (such as securities) in the underlying transaction.
	Quantity *FinancialInstrumentQuantityChoice `xml:"Qty"`

	// Face amount and amortised value of security.
	OriginalAndCurrentFaceAmount *OriginalAndCurrentQuantities1 `xml:"OrgnlAndCurFaceAmt"`

	// Proprietary quantities specification defined in the underlying transaction.
	Proprietary *ProprietaryQuantity1 `xml:"Prtry"`
}

func (t *TransactionQuantities2Choice) AddQuantity() *FinancialInstrumentQuantityChoice {
	t.Quantity = new(FinancialInstrumentQuantityChoice)
	return t.Quantity
}

func (t *TransactionQuantities2Choice) AddOriginalAndCurrentFaceAmount() *OriginalAndCurrentQuantities1 {
	t.OriginalAndCurrentFaceAmount = new(OriginalAndCurrentQuantities1)
	return t.OriginalAndCurrentFaceAmount
}

func (t *TransactionQuantities2Choice) AddProprietary() *ProprietaryQuantity1 {
	t.Proprietary = new(ProprietaryQuantity1)
	return t.Proprietary
}
