package model

// Information used to calculate the tax.
type TaxCalculationInformation5 struct {

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	Basis *TaxationBasis2Code `xml:"Bsis,omitempty"`

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	ExtendedBasis *Extended350Code `xml:"XtndedBsis,omitempty"`
}

func (t *TaxCalculationInformation5) SetBasis(value string) {
	t.Basis = (*TaxationBasis2Code)(&value)
}

func (t *TaxCalculationInformation5) SetExtendedBasis(value string) {
	t.ExtendedBasis = (*Extended350Code)(&value)
}
