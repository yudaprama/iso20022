package model

// Information regarding the total amount of taxes.
type TotalTaxes2 struct {

	// Total value of the taxes for a specific order.
	TotalAmountOfTaxes *ActiveCurrencyAnd13DecimalAmount `xml:"TtlAmtOfTaxs,omitempty"`

	// Information related to a specific tax.
	TaxDetails []*Tax7 `xml:"TaxDtls"`
}

func (t *TotalTaxes2) SetTotalAmountOfTaxes(value, currency string) {
	t.TotalAmountOfTaxes = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TotalTaxes2) AddTaxDetails() *Tax7 {
	newValue := new(Tax7)
	t.TaxDetails = append(t.TaxDetails, newValue)
	return newValue
}
