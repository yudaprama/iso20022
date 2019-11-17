package model

// Resulting debit or credit amount of the netted amounts for all debit and credit entries.
type AmountAndDirection35 struct {

	// Resulting amount of the netted amounts for all debit and credit entries.
	Amount *NonNegativeDecimalNumber `xml:"Amt"`

	// Indicates whether the amount is a credit or a debit amount.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`
}

func (a *AmountAndDirection35) SetAmount(value string) {
	a.Amount = (*NonNegativeDecimalNumber)(&value)
}

func (a *AmountAndDirection35) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}
