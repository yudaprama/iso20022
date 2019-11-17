package model

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount1 struct {

	// Type of amount.
	Type *TypeOfAmount2Code `xml:"Tp"`

	// Amount value.
	Value *ImpliedCurrencyAndAmount `xml:"Val"`
}

func (d *DetailedAmount1) SetType(value string) {
	d.Type = (*TypeOfAmount2Code)(&value)
}

func (d *DetailedAmount1) SetValue(value, currency string) {
	d.Value = NewImpliedCurrencyAndAmount(value, currency)
}
