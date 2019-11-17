package model

// Account to or from which a cash entry is made.
type CashAccount12 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *CashAccountIdentification1Choice `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm"`

	// Specifies the nature, or use, of the cash account.
	Type *CashAccountType1Code `xml:"Tp,omitempty"`

	// Specifies the nature, or use, of the cash account.
	ExtendedType *Extended350Code `xml:"XtndedTp,omitempty"`

	// Medium of exchange of value.
	Currency *ActiveCurrencyCode `xml:"Ccy"`

	// Specifies the current state of an account, eg, enabled or deleted.
	Status *AccountStatus1Code `xml:"Sts"`
}

func (c *CashAccount12) AddIdentification() *CashAccountIdentification1Choice {
	c.Identification = new(CashAccountIdentification1Choice)
	return c.Identification
}

func (c *CashAccount12) SetName(value string) {
	c.Name = (*Max35Text)(&value)
}

func (c *CashAccount12) SetType(value string) {
	c.Type = (*CashAccountType1Code)(&value)
}

func (c *CashAccount12) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *CashAccount12) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CashAccount12) SetStatus(value string) {
	c.Status = (*AccountStatus1Code)(&value)
}
