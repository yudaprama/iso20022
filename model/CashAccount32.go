package model

// Account to or from which a cash entry is made.
type CashAccount32 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationAndName5 `xml:"Id"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification70Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	AccountServicerBranch *BranchData `xml:"AcctSvcrBrnch,omitempty"`
}

func (c *CashAccount32) AddIdentification() *AccountIdentificationAndName5 {
	c.Identification = new(AccountIdentificationAndName5)
	return c.Identification
}

func (c *CashAccount32) AddAccountOwner() *PartyIdentification70Choice {
	c.AccountOwner = new(PartyIdentification70Choice)
	return c.AccountOwner
}

func (c *CashAccount32) AddAccountServicer() *PartyIdentification70Choice {
	c.AccountServicer = new(PartyIdentification70Choice)
	return c.AccountServicer
}

func (c *CashAccount32) AddAccountServicerBranch() *BranchData {
	c.AccountServicerBranch = new(BranchData)
	return c.AccountServicerBranch
}
