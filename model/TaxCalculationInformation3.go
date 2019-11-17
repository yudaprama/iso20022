package model

// Information used to calculate the tax.
type TaxCalculationInformation3 struct {

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	Basis *TaxationBasis1 `xml:"Bsis,omitempty"`

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUCapitalGain *EUCapitalGain1 `xml:"EUCptlGn,omitempty"`

	// Amount of money that it is to be taxed.
	TaxableAmount *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblAmt,omitempty"`
}

func (t *TaxCalculationInformation3) AddBasis() *TaxationBasis1 {
	t.Basis = new(TaxationBasis1)
	return t.Basis
}

func (t *TaxCalculationInformation3) AddEUCapitalGain() *EUCapitalGain1 {
	t.EUCapitalGain = new(EUCapitalGain1)
	return t.EUCapitalGain
}

func (t *TaxCalculationInformation3) SetTaxableAmount(value, currency string) {
	t.TaxableAmount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
