package model

// Tax related to an investment fund order.
type Tax8 struct {

	// Type of tax applied.
	Type *TaxType3 `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation2 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax8) AddType() *TaxType3 {
	t.Type = new(TaxType3)
	return t.Type
}

func (t *Tax8) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax8) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *Tax8) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax8) AddTaxCalculationDetails() *TaxCalculationInformation2 {
	t.TaxCalculationDetails = new(TaxCalculationInformation2)
	return t.TaxCalculationDetails
}
