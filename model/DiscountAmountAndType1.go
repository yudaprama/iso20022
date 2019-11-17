package model

// Specifies the amount with a specific type.
type DiscountAmountAndType1 struct {

	// Specifies the type of the amount.
	Type *DiscountAmountType1Choice `xml:"Tp,omitempty"`

	// Amount of money, which has been typed.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (d *DiscountAmountAndType1) AddType() *DiscountAmountType1Choice {
	d.Type = new(DiscountAmountType1Choice)
	return d.Type
}

func (d *DiscountAmountAndType1) SetAmount(value, currency string) {
	d.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
