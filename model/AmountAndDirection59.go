package model

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection59 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`
}

func (a *AmountAndDirection59) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection59) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}
