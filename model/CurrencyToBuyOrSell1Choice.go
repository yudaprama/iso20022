package model

// Choice of counterparty type.
type CurrencyToBuyOrSell1Choice struct {

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell"`
}

func (c *CurrencyToBuyOrSell1Choice) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CurrencyToBuyOrSell1Choice) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}
