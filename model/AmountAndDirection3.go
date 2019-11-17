package model

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection3 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebit *CreditDebitCode `xml:"CdtDbt"`
}

func (a *AmountAndDirection3) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection3) SetCreditDebit(value string) {
	a.CreditDebit = (*CreditDebitCode)(&value)
}
