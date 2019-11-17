package model

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection18 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebit *CreditDebitCode `xml:"CdtDbt"`
}

func (a *AmountAndDirection18) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection18) SetCreditDebit(value string) {
	a.CreditDebit = (*CreditDebitCode)(&value)
}
