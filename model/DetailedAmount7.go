package model

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount7 struct {

	// Cash-back amount.
	CashBack *ImpliedCurrencyAndAmount `xml:"CshBck,omitempty"`

	// Gratuity amount.
	Gratuity *ImpliedCurrencyAndAmount `xml:"Grtty,omitempty"`

	// Fees amount.
	Fees []*DetailedAmount4 `xml:"Fees,omitempty"`

	// Global rebate of the transaction. This amount is counted as a negative amount.
	Rebate []*DetailedAmount4 `xml:"Rbt,omitempty"`

	// Value added tax amount.
	ValueAddedTax []*DetailedAmount4 `xml:"ValAddedTax,omitempty"`

	// Additional charge paid by the cardholder. For example airline credit card surcharge.
	Surcharge []*DetailedAmount4 `xml:"Srchrg,omitempty"`
}

func (d *DetailedAmount7) SetCashBack(value, currency string) {
	d.CashBack = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount7) SetGratuity(value, currency string) {
	d.Gratuity = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount7) AddFees() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.Fees = append(d.Fees, newValue)
	return newValue
}

func (d *DetailedAmount7) AddRebate() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.Rebate = append(d.Rebate, newValue)
	return newValue
}

func (d *DetailedAmount7) AddValueAddedTax() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.ValueAddedTax = append(d.ValueAddedTax, newValue)
	return newValue
}

func (d *DetailedAmount7) AddSurcharge() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.Surcharge = append(d.Surcharge, newValue)
	return newValue
}
