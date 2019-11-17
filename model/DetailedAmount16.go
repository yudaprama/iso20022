package model

// Amounts of the deposit transaction.
type DetailedAmount16 struct {

	// Link to the account for multi-account deposit.
	AccountSequenceNumber *Number `xml:"AcctSeqNb,omitempty"`

	// Amount of the deposit to be made on the ATM after the approval of the deposit transaction.
	AmountToDeposit *ImpliedCurrencyAndAmount `xml:"AmtToDpst,omitempty"`

	// Currency of the amount to deposit when different from the base currency of the ATM.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Cashback amount value.
	CashBackAmount *ImpliedCurrencyAndAmount `xml:"CshBckAmt,omitempty"`

	// Deposit fees, accepted by the customer.
	Fees []*DetailedAmount13 `xml:"Fees,omitempty"`

	// Amount of the donation.
	Donation []*DetailedAmount13 `xml:"Dontn,omitempty"`
}

func (d *DetailedAmount16) SetAccountSequenceNumber(value string) {
	d.AccountSequenceNumber = (*Number)(&value)
}

func (d *DetailedAmount16) SetAmountToDeposit(value, currency string) {
	d.AmountToDeposit = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount16) SetCurrency(value string) {
	d.Currency = (*ActiveCurrencyCode)(&value)
}

func (d *DetailedAmount16) SetCashBackAmount(value, currency string) {
	d.CashBackAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount16) AddFees() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	d.Fees = append(d.Fees, newValue)
	return newValue
}

func (d *DetailedAmount16) AddDonation() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	d.Donation = append(d.Donation, newValue)
	return newValue
}
