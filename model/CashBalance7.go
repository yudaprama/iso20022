package model

// Set of elements used to define the balance details.
type CashBalance7 struct {

	// Specifies the nature of a balance.
	Type *BalanceType12 `xml:"Tp"`

	// Set of elements used to provide details on the credit line.
	CreditLine *CreditLine2 `xml:"CdtLine,omitempty"`

	// Amount of money of the cash balance.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the balance is a credit or a debit balance.
	// Usage: A zero balance is considered to be a credit balance.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates the date (and time) of the balance.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Set of elements used to indicate when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashAvailability1 `xml:"Avlbty,omitempty"`
}

func (c *CashBalance7) AddType() *BalanceType12 {
	c.Type = new(BalanceType12)
	return c.Type
}

func (c *CashBalance7) AddCreditLine() *CreditLine2 {
	c.CreditLine = new(CreditLine2)
	return c.CreditLine
}

func (c *CashBalance7) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashBalance7) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashBalance7) AddDate() *DateAndDateTimeChoice {
	c.Date = new(DateAndDateTimeChoice)
	return c.Date
}

func (c *CashBalance7) AddAvailability() *CashAvailability1 {
	newValue := new(CashAvailability1)
	c.Availability = append(c.Availability, newValue)
	return newValue
}
