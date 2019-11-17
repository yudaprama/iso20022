package model

// Provides information about the cash movement.
type CashMovement2 struct {

	// Cash amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Provides information about the account which is debited/credited.
	AccountDetails []*CashAccount19 `xml:"AcctDtls"`
}

func (c *CashMovement2) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashMovement2) AddAccountDetails() *CashAccount19 {
	newValue := new(CashAccount19)
	c.AccountDetails = append(c.AccountDetails, newValue)
	return newValue
}
