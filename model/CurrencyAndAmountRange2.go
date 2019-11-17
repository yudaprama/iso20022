package model

// Range of amount values.
type CurrencyAndAmountRange2 struct {

	// Specified amount or amount range.
	Amount *ImpliedCurrencyAmountRangeChoice `xml:"Amt"`

	// Indicates whether the amount is a credited or debited amount.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Medium of exchange of value, used to qualify an amount.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy"`
}

func (c *CurrencyAndAmountRange2) AddAmount() *ImpliedCurrencyAmountRangeChoice {
	c.Amount = new(ImpliedCurrencyAmountRangeChoice)
	return c.Amount
}

func (c *CurrencyAndAmountRange2) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CurrencyAndAmountRange2) SetCurrency(value string) {
	c.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}
