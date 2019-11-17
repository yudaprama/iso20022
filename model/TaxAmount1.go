package model

// Set of elements used to provide information on the tax amount(s) of tax record.
type TaxAmount1 struct {

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Amount of money on which the tax is based.
	TaxableBaseAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`

	// Total amount that is the result of the calculation of the tax for the record.
	TotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Set of elements used to provide details on the tax period and amount.
	Details []*TaxRecordDetails1 `xml:"Dtls,omitempty"`
}

func (t *TaxAmount1) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxAmount1) SetTaxableBaseAmount(value, currency string) {
	t.TaxableBaseAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxAmount1) SetTotalAmount(value, currency string) {
	t.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxAmount1) AddDetails() *TaxRecordDetails1 {
	newValue := new(TaxRecordDetails1)
	t.Details = append(t.Details, newValue)
	return newValue
}
