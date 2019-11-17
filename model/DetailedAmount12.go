package model

// Amounts of the withdrawal transaction.
type DetailedAmount12 struct {

	// Amount to be dispensed by the ATM after the approval of the withdrawal transaction.
	AmountToDispense *ImpliedCurrencyAndAmount `xml:"AmtToDspns"`

	// Currency of the amount to dispense when different from the base currency of the ATM.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Withdrawal fees, accepted by the customer.
	Fees []*DetailedAmount13 `xml:"Fees,omitempty"`

	// Amount of the donation.
	Donation []*DetailedAmount13 `xml:"Dontn,omitempty"`
}

func (d *DetailedAmount12) SetAmountToDispense(value, currency string) {
	d.AmountToDispense = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount12) SetCurrency(value string) {
	d.Currency = (*ActiveCurrencyCode)(&value)
}

func (d *DetailedAmount12) AddFees() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	d.Fees = append(d.Fees, newValue)
	return newValue
}

func (d *DetailedAmount12) AddDonation() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	d.Donation = append(d.Donation, newValue)
	return newValue
}
