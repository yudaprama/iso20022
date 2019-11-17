package model

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection23 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection23) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection23) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection23) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return a.ForeignExchangeDetails
}
