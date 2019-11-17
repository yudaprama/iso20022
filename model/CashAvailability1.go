package model

// Indicates when the amount of money will become available, that is can be accessed and start generating interest.
type CashAvailability1 struct {

	// Indicates when the amount of money will become available.
	Date *CashAvailabilityDate1Choice `xml:"Dt"`

	// Identifies the available amount.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the availability balance is a credit or a debit balance.
	// Usage: A zero balance is considered to be a credit balance.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`
}

func (c *CashAvailability1) AddDate() *CashAvailabilityDate1Choice {
	c.Date = new(CashAvailabilityDate1Choice)
	return c.Date
}

func (c *CashAvailability1) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashAvailability1) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}
