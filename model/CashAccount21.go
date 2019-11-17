package model

// Account to or from which a cash entry is made.
type CashAccount21 struct {

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *BICIdentifier `xml:"Svcr,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification5Choice `xml:"Id"`
}

func (c *CashAccount21) SetServicer(value string) {
	c.Servicer = (*BICIdentifier)(&value)
}

func (c *CashAccount21) AddIdentification() *AccountIdentification5Choice {
	c.Identification = new(AccountIdentification5Choice)
	return c.Identification
}
