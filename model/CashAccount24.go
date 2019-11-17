package model

// Provides the details to identify an account.
type CashAccount24 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4Choice `xml:"Id"`

	// Specifies the nature, or use of the account.
	Type *CashAccountType2Choice `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	//
	// Usage: Currency should only be used in case one and the same account number covers several currencies
	// and the initiating party needs to identify which currency needs to be used for settlement on the account.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (c *CashAccount24) AddIdentification() *AccountIdentification4Choice {
	c.Identification = new(AccountIdentification4Choice)
	return c.Identification
}

func (c *CashAccount24) AddType() *CashAccountType2Choice {
	c.Type = new(CashAccountType2Choice)
	return c.Type
}

func (c *CashAccount24) SetCurrency(value string) {
	c.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccount24) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}
