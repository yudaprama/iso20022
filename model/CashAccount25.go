package model

// Provides the details to identify an account.
type CashAccount25 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4Choice `xml:"Id"`

	// Specifies the nature, or use of the account.
	Type *CashAccountType2Choice `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Party that legally owns the account.
	Owner *PartyIdentification43 `xml:"Ownr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *BranchAndFinancialInstitutionIdentification5 `xml:"Svcr,omitempty"`
}

func (c *CashAccount25) AddIdentification() *AccountIdentification4Choice {
	c.Identification = new(AccountIdentification4Choice)
	return c.Identification
}

func (c *CashAccount25) AddType() *CashAccountType2Choice {
	c.Type = new(CashAccountType2Choice)
	return c.Type
}

func (c *CashAccount25) SetCurrency(value string) {
	c.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccount25) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

func (c *CashAccount25) AddOwner() *PartyIdentification43 {
	c.Owner = new(PartyIdentification43)
	return c.Owner
}

func (c *CashAccount25) AddServicer() *BranchAndFinancialInstitutionIdentification5 {
	c.Servicer = new(BranchAndFinancialInstitutionIdentification5)
	return c.Servicer
}
