package model

// Specifies the amount with a specific type.
type TaxAmountAndType1 struct {

	// Specifies the type of the amount.
	Type *TaxAmountType1Choice `xml:"Tp,omitempty"`

	// Amount of money, which has been typed.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (t *TaxAmountAndType1) AddType() *TaxAmountType1Choice {
	t.Type = new(TaxAmountType1Choice)
	return t.Type
}

func (t *TaxAmountAndType1) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
