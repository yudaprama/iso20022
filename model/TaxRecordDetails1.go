package model

// Provides information on the individual tax amount(s) per period of the tax record.
type TaxRecordDetails1 struct {

	// Set of elements used to provide details on the period of time related to the tax payment.
	Period *TaxPeriod1 `xml:"Prd,omitempty"`

	// Underlying tax amount related to the specified period.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (t *TaxRecordDetails1) AddPeriod() *TaxPeriod1 {
	t.Period = new(TaxPeriod1)
	return t.Period
}

func (t *TaxRecordDetails1) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
