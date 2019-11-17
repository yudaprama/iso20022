package model

// Information used to calculate the tax.
type TaxCalculationInformation6 struct {

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	Basis *TaxationBasis2Code `xml:"Bsis,omitempty"`

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	ExtendedBasis *Extended350Code `xml:"XtndedBsis,omitempty"`

	// Amount of money that it is to be taxed.
	TaxableAmount *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblAmt,omitempty"`
}

func (t *TaxCalculationInformation6) SetBasis(value string) {
	t.Basis = (*TaxationBasis2Code)(&value)
}

func (t *TaxCalculationInformation6) SetExtendedBasis(value string) {
	t.ExtendedBasis = (*Extended350Code)(&value)
}

func (t *TaxCalculationInformation6) SetTaxableAmount(value, currency string) {
	t.TaxableAmount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
