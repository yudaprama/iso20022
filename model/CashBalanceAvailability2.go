package model

// Indicates when the amount of money will become available, ie can be accessed and start generating interest.
type CashBalanceAvailability2 struct {

	// Indicates when the amount of money will become available.
	Date *CashBalanceAvailabilityDate1 `xml:"Dt"`

	// Identifies the available amount.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the availability balance is a credit or a debit balance.
	// Usage: A zero balance is considered to be a credit balance.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`
}

func (c *CashBalanceAvailability2) AddDate() *CashBalanceAvailabilityDate1 {
	c.Date = new(CashBalanceAvailabilityDate1)
	return c.Date
}

func (c *CashBalanceAvailability2) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashBalanceAvailability2) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}
