package model

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection44 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection44) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection44) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection44) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection44) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return a.ForeignExchangeDetails
}
