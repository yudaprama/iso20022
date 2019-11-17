package model

// Details of the transfer transaction amounts.
type DetailedAmount17 struct {

	// Amount to be transferred from the source account to the destination account.
	AmountToTransfer *ImpliedCurrencyAndAmount `xml:"AmtToTrf"`

	// Currency of the amount to be transferred.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Transfer fees, accepted by the customer.
	Fees []*DetailedAmount18 `xml:"Fees,omitempty"`

	// Amount of the donation.
	Donation []*DetailedAmount18 `xml:"Dontn,omitempty"`
}

func (d *DetailedAmount17) SetAmountToTransfer(value, currency string) {
	d.AmountToTransfer = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount17) SetCurrency(value string) {
	d.Currency = (*ActiveCurrencyCode)(&value)
}

func (d *DetailedAmount17) AddFees() *DetailedAmount18 {
	newValue := new(DetailedAmount18)
	d.Fees = append(d.Fees, newValue)
	return newValue
}

func (d *DetailedAmount17) AddDonation() *DetailedAmount18 {
	newValue := new(DetailedAmount18)
	d.Donation = append(d.Donation, newValue)
	return newValue
}
