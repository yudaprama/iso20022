package model

// Information used to calculate the tax.
type TaxCalculationInformation1 struct {

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	Basis *TaxationBasis1 `xml:"Bsis,omitempty"`
}

func (t *TaxCalculationInformation1) AddBasis() *TaxationBasis1 {
	t.Basis = new(TaxationBasis1)
	return t.Basis
}
