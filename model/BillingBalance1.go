package model

// Provides the balance for the billing services.
type BillingBalance1 struct {

	// Defines the type of  balance.
	Type *BillingBalanceType1Choice `xml:"Tp"`

	// Balance value.
	Value *AmountAndDirection34 `xml:"Val"`

	// Identifies the currency type used to report the balance. This is not the ISO code.
	CurrencyType *BillingCurrencyType1Code `xml:"CcyTp,omitempty"`
}

func (b *BillingBalance1) AddType() *BillingBalanceType1Choice {
	b.Type = new(BillingBalanceType1Choice)
	return b.Type
}

func (b *BillingBalance1) AddValue() *AmountAndDirection34 {
	b.Value = new(AmountAndDirection34)
	return b.Value
}

func (b *BillingBalance1) SetCurrencyType(value string) {
	b.CurrencyType = (*BillingCurrencyType1Code)(&value)
}
