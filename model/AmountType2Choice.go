package model

// Amount of money to be moved between the debtor and creditor, expressed in debtor's account currency or converted in another currency.
type AmountType2Choice struct {

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *CurrencyAndAmount `xml:"InstdAmt"`

	// Amount of money to be transferred between the debtor and creditor, before deduction of charges, expressed in the currency of the debtor's account, and to be transferred into a different currency.
	//
	// Usage : Currency of the amount is expressed in the currency of the debtor's account, but the amount to be transferred is in another currency. The debtor agent will convert the amount and currency to the to be transferred amount and currency, eg, 'pay equivalent of 100000 EUR in JPY'(and account is in EUR).
	EquivalentAmount *EquivalentAmount `xml:"EqvtAmt"`
}

func (a *AmountType2Choice) SetInstructedAmount(value, currency string) {
	a.InstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (a *AmountType2Choice) AddEquivalentAmount() *EquivalentAmount {
	a.EquivalentAmount = new(EquivalentAmount)
	return a.EquivalentAmount
}
