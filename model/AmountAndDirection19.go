package model

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection19 struct {

	// Amount of money that is debited or credited.
	Amount *RestrictedFINActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates if the amount is a debited or a credited.
	CreditDebit *CreditDebitCode `xml:"CdtDbt,omitempty"`
}

func (a *AmountAndDirection19) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection19) SetCreditDebit(value string) {
	a.CreditDebit = (*CreditDebitCode)(&value)
}
