package model

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection72 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms27 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection72) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection72) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection72) AddForeignExchangeDetails() *ForeignExchangeTerms27 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms27)
	return a.ForeignExchangeDetails
}
