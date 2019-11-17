package model

// Information used to calculate the tax.
type TaxCalculationInformation10 struct {

	// Form of the rebate, for example, cash.
	Basis *TaxBasis1Choice `xml:"Bsis,omitempty"`

	// Amount of money on which the tax is charged.
	TaxableAmount *ActiveCurrencyAndAmount `xml:"TaxblAmt"`
}

func (t *TaxCalculationInformation10) AddBasis() *TaxBasis1Choice {
	t.Basis = new(TaxBasis1Choice)
	return t.Basis
}

func (t *TaxCalculationInformation10) SetTaxableAmount(value, currency string) {
	t.TaxableAmount = NewActiveCurrencyAndAmount(value, currency)
}
