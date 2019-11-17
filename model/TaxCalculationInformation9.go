package model

// Information used to calculate the tax.
type TaxCalculationInformation9 struct {

	// Form of the rebate, for example, cash.
	Basis *TaxBasis1Choice `xml:"Bsis"`
}

func (t *TaxCalculationInformation9) AddBasis() *TaxBasis1Choice {
	t.Basis = new(TaxBasis1Choice)
	return t.Basis
}
