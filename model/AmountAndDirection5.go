package model

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection5 struct {

	// Amount of money that is debited or credited.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates if the amount is a debited or a credited.
	CreditDebit *CreditDebitCode `xml:"CdtDbt,omitempty"`
}

func (a *AmountAndDirection5) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection5) SetCreditDebit(value string) {
	a.CreditDebit = (*CreditDebitCode)(&value)
}
