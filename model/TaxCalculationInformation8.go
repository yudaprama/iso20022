package model

// Information used to calculate the tax.
type TaxCalculationInformation8 struct {

	// Form of the rebate, for example, cash.
	Basis *TaxBasis1Choice `xml:"Bsis,omitempty"`

	// Amount of money on which the tax is charged.
	TaxableAmount *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblAmt"`
}

func (t *TaxCalculationInformation8) AddBasis() *TaxBasis1Choice {
	t.Basis = new(TaxBasis1Choice)
	return t.Basis
}

func (t *TaxCalculationInformation8) SetTaxableAmount(value, currency string) {
	t.TaxableAmount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
