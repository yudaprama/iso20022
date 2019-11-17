package model

// Set of elements used to indicate when the booked amount of money will become available, ie can be accessed and start generating interest.
type CashBalanceAvailability1 struct {

	// Indicates when the amount of money will become available.
	Date *CashBalanceAvailabilityDate1 `xml:"Dt"`

	// Identifies the available amount.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Indicates whether the availability balance is a credit or a debit balance. A zero balance is considered to be a credit balance
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`
}

func (c *CashBalanceAvailability1) AddDate() *CashBalanceAvailabilityDate1 {
	c.Date = new(CashBalanceAvailabilityDate1)
	return c.Date
}

func (c *CashBalanceAvailability1) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}

func (c *CashBalanceAvailability1) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}
