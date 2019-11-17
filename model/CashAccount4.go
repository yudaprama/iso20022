package model

// Account to or from which a cash entry is made.
type CashAccount4 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationAndName3 `xml:"Id"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification2Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	AccountServicerBranch *BranchData `xml:"AcctSvcrBrnch,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	InvestmentAccountType *FundCashAccount2Code `xml:"InvstmtAcctTp,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	ExtendedInvestmentAccountType *Extended350Code `xml:"XtndedInvstmtAcctTp,omitempty"`
}

func (c *CashAccount4) AddIdentification() *AccountIdentificationAndName3 {
	c.Identification = new(AccountIdentificationAndName3)
	return c.Identification
}

func (c *CashAccount4) AddAccountOwner() *PartyIdentification2Choice {
	c.AccountOwner = new(PartyIdentification2Choice)
	return c.AccountOwner
}

func (c *CashAccount4) AddAccountServicer() *PartyIdentification2Choice {
	c.AccountServicer = new(PartyIdentification2Choice)
	return c.AccountServicer
}

func (c *CashAccount4) AddAccountServicerBranch() *BranchData {
	c.AccountServicerBranch = new(BranchData)
	return c.AccountServicerBranch
}

func (c *CashAccount4) SetInvestmentAccountType(value string) {
	c.InvestmentAccountType = (*FundCashAccount2Code)(&value)
}

func (c *CashAccount4) SetExtendedInvestmentAccountType(value string) {
	c.ExtendedInvestmentAccountType = (*Extended350Code)(&value)
}
