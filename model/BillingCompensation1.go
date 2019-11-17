package model

// Specifies the compensation data of an incorrect billing.
type BillingCompensation1 struct {

	// Defines the type of billing compensation.
	Type *BillingCompensationType1Choice `xml:"Tp"`

	// Defines the value of compensation.
	Value *AmountAndDirection34 `xml:"Val"`

	// Identifies the currency type used to report the value or total, in a coded form, such as Settlement (STLM).
	CurrencyType *BillingCurrencyType2Code `xml:"CcyTp,omitempty"`
}

func (b *BillingCompensation1) AddType() *BillingCompensationType1Choice {
	b.Type = new(BillingCompensationType1Choice)
	return b.Type
}

func (b *BillingCompensation1) AddValue() *AmountAndDirection34 {
	b.Value = new(AmountAndDirection34)
	return b.Value
}

func (b *BillingCompensation1) SetCurrencyType(value string) {
	b.CurrencyType = (*BillingCurrencyType2Code)(&value)
}
