package model

// Information used for identifying an account.
type CashAccount13 struct {

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Identification *AccountIdentification3Choice `xml:"Id"`

	// Nature, or use, of the account.
	Type *CashAccountType2 `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, assigned by the account servicing institution in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage : the account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Party that legally owns the account.
	Owner *PartyIdentification8 `xml:"Ownr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *BranchAndFinancialInstitutionIdentification3 `xml:"Svcr,omitempty"`
}

func (c *CashAccount13) AddIdentification() *AccountIdentification3Choice {
	c.Identification = new(AccountIdentification3Choice)
	return c.Identification
}

func (c *CashAccount13) AddType() *CashAccountType2 {
	c.Type = new(CashAccountType2)
	return c.Type
}

func (c *CashAccount13) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CashAccount13) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

func (c *CashAccount13) AddOwner() *PartyIdentification8 {
	c.Owner = new(PartyIdentification8)
	return c.Owner
}

func (c *CashAccount13) AddServicer() *BranchAndFinancialInstitutionIdentification3 {
	c.Servicer = new(BranchAndFinancialInstitutionIdentification3)
	return c.Servicer
}
