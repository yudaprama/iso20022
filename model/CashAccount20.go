package model

// Set of elements used to identify an account.
type CashAccount20 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4Choice `xml:"Id"`

	// Specifies the nature, or use of the account.
	Type *CashAccountType2 `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Party that legally owns the account.
	Owner *PartyIdentification32 `xml:"Ownr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *BranchAndFinancialInstitutionIdentification4 `xml:"Svcr,omitempty"`
}

func (c *CashAccount20) AddIdentification() *AccountIdentification4Choice {
	c.Identification = new(AccountIdentification4Choice)
	return c.Identification
}

func (c *CashAccount20) AddType() *CashAccountType2 {
	c.Type = new(CashAccountType2)
	return c.Type
}

func (c *CashAccount20) SetCurrency(value string) {
	c.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccount20) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

func (c *CashAccount20) AddOwner() *PartyIdentification32 {
	c.Owner = new(PartyIdentification32)
	return c.Owner
}

func (c *CashAccount20) AddServicer() *BranchAndFinancialInstitutionIdentification4 {
	c.Servicer = new(BranchAndFinancialInstitutionIdentification4)
	return c.Servicer
}
