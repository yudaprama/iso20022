package model

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection47 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection47) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection47) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection47) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return a.ForeignExchangeDetails
}
