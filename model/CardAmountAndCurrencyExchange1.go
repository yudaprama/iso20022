package model

// Provides details on the detailed or original amount and currency.
type CardAmountAndCurrencyExchange1 struct {

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Reports on currency exchange information.
	CurrencyExchange *CurrencyExchange3 `xml:"CcyXchg,omitempty"`

	// Identification or qualification of the type of amount.
	Type *TypeOfAmount3Code `xml:"Tp,omitempty"`
}

func (c *CardAmountAndCurrencyExchange1) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CardAmountAndCurrencyExchange1) AddCurrencyExchange() *CurrencyExchange3 {
	c.CurrencyExchange = new(CurrencyExchange3)
	return c.CurrencyExchange
}

func (c *CardAmountAndCurrencyExchange1) SetType(value string) {
	c.Type = (*TypeOfAmount3Code)(&value)
}
