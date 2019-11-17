package model

// Information used for identifying an account.
type CashAccount3 struct {

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Identification *AccountIdentification1Choice `xml:"Id"`

	// Nature, or use, of the account.
	Type *CashAccountType3Code `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	//
	// Usage: Currency should only be used in case one and the same account number covers several currencies and the initiating party needs to identify which currency needs to be used for settlement on the account.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, assigned by the account servicing institution in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage : the account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (c *CashAccount3) AddIdentification() *AccountIdentification1Choice {
	c.Identification = new(AccountIdentification1Choice)
	return c.Identification
}

func (c *CashAccount3) SetType(value string) {
	c.Type = (*CashAccountType3Code)(&value)
}

func (c *CashAccount3) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CashAccount3) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}
