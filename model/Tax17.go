package model

// Tax related to an investment fund order.
type Tax17 struct {

	// Type of tax applied.
	Type *TaxType12Code `xml:"Tp"`

	// Type of tax applied.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money resulting from the calculation of the tax.
	Amount []*ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation4 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax17) SetType(value string) {
	t.Type = (*TaxType12Code)(&value)
}

func (t *Tax17) SetExtendedType(value string) {
	t.ExtendedType = (*Extended350Code)(&value)
}

func (t *Tax17) AddAmount(value, currency string) {
	t.Amount = append(t.Amount, NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency))
}

func (t *Tax17) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *Tax17) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax17) AddTaxCalculationDetails() *TaxCalculationInformation4 {
	t.TaxCalculationDetails = new(TaxCalculationInformation4)
	return t.TaxCalculationDetails
}
