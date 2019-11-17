package model

// Additional amounts from the processor or the issuer without financial impacts on the transaction amount.
type DetailedAmount10 struct {

	// Type or class of amount.
	Type *TypeOfAmount6Code `xml:"Tp"`

	// Additional information to specify the type of amount.
	AdditionalType *Max35Text `xml:"AddtlTp,omitempty"`

	// Amount value.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Short description of the amount to provide to the cardholder.
	Label *Max35Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount10) SetType(value string) {
	d.Type = (*TypeOfAmount6Code)(&value)
}

func (d *DetailedAmount10) SetAdditionalType(value string) {
	d.AdditionalType = (*Max35Text)(&value)
}

func (d *DetailedAmount10) SetAmount(value, currency string) {
	d.Amount = NewCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount10) SetLabel(value string) {
	d.Label = (*Max35Text)(&value)
}
