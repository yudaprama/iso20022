package model

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type Tax22 struct {

	// Specifies the type of tax.
	Type *TaxType2Choice `xml:"Tp"`

	// Specifies the tax as an amount.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (t *Tax22) AddType() *TaxType2Choice {
	t.Type = new(TaxType2Choice)
	return t.Type
}

func (t *Tax22) SetAmount(value, currency string) {
	t.Amount = NewCurrencyAndAmount(value, currency)
}
