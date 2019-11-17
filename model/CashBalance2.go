package model

// Set of elements defining the balance details.
type CashBalance2 struct {

	// Specifies the nature of a balance, eg, opening booked balance.
	Type *BalanceType2Choice `xml:"Tp"`

	// Provides further details on the credit line information.
	CreditLine *CreditLine1 `xml:"CdtLine,omitempty"`

	// Currency and amount of money of the cash balance.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Indicates whether the balance is a credit or a debit balance. A zero balance is considered to be a credit balance
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies the date (and time) of the balance.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Set of elements used to indicate when the booked funds will become available, ie can be accessed and start generating interest.
	//
	// Usage : this type of info is eg used in US, and is linked to particular instruments, such as cheques.
	// Example : When a cheque is deposited, it will be booked on the deposit day, but the funds will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability1 `xml:"Avlbty,omitempty"`
}

func (c *CashBalance2) AddType() *BalanceType2Choice {
	c.Type = new(BalanceType2Choice)
	return c.Type
}

func (c *CashBalance2) AddCreditLine() *CreditLine1 {
	c.CreditLine = new(CreditLine1)
	return c.CreditLine
}

func (c *CashBalance2) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}

func (c *CashBalance2) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashBalance2) AddDate() *DateAndDateTimeChoice {
	c.Date = new(DateAndDateTimeChoice)
	return c.Date
}

func (c *CashBalance2) AddAvailability() *CashBalanceAvailability1 {
	newValue := new(CashBalanceAvailability1)
	c.Availability = append(c.Availability, newValue)
	return newValue
}
