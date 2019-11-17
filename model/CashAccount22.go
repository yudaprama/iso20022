package model

// Account to or from which a cash entry is made.
type CashAccount22 struct {

	// Medium of exchange of value.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *BICIdentifier `xml:"Svcr"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification5Choice `xml:"Id"`

	// Sub-division of a master or omnibus cash account.
	SecondaryAccount *CashAccount21 `xml:"ScndryAcct,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountTypeDescription *Max35Text `xml:"AcctTpDesc"`
}

func (c *CashAccount22) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CashAccount22) SetServicer(value string) {
	c.Servicer = (*BICIdentifier)(&value)
}

func (c *CashAccount22) AddIdentification() *AccountIdentification5Choice {
	c.Identification = new(AccountIdentification5Choice)
	return c.Identification
}

func (c *CashAccount22) AddSecondaryAccount() *CashAccount21 {
	c.SecondaryAccount = new(CashAccount21)
	return c.SecondaryAccount
}

func (c *CashAccount22) SetAccountTypeDescription(value string) {
	c.AccountTypeDescription = (*Max35Text)(&value)
}
