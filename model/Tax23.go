package model

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type Tax23 struct {

	// Specifies the type of tax.
	Type *TaxType2Choice `xml:"Tp"`

	// Specifies the tax as an amount or percentage.
	AmountOrPercentage *AmountOrPercentage2Choice `xml:"AmtOrPctg"`
}

func (t *Tax23) AddType() *TaxType2Choice {
	t.Type = new(TaxType2Choice)
	return t.Type
}

func (t *Tax23) AddAmountOrPercentage() *AmountOrPercentage2Choice {
	t.AmountOrPercentage = new(AmountOrPercentage2Choice)
	return t.AmountOrPercentage
}
