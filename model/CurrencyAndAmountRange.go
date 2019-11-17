package model

// Range of amount values.
type CurrencyAndAmountRange struct {

	// Specified amount or amount range.
	Amount *ImpliedCurrencyAmountRangeChoice `xml:"Amt"`

	// Indicates whether the amount is a credited or debited amount.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Medium of exchange of value, used to qualify an amount.
	Currency *CurrencyCode `xml:"Ccy"`
}

func (c *CurrencyAndAmountRange) AddAmount() *ImpliedCurrencyAmountRangeChoice {
	c.Amount = new(ImpliedCurrencyAmountRangeChoice)
	return c.Amount
}

func (c *CurrencyAndAmountRange) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CurrencyAndAmountRange) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}
