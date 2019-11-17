package model

// Specifies the amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
type AmountType4Choice struct {

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	//
	// Usage: This amount has to be transported unchanged through the transaction chain.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`

	// Amount of money to be moved between the debtor and creditor, expressed in the currency of the debtor's account, and the currency in which the amount is to be moved.
	EquivalentAmount *EquivalentAmount2 `xml:"EqvtAmt"`
}

func (a *AmountType4Choice) SetInstructedAmount(value, currency string) {
	a.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountType4Choice) AddEquivalentAmount() *EquivalentAmount2 {
	a.EquivalentAmount = new(EquivalentAmount2)
	return a.EquivalentAmount
}
