package model

// Information used for identifying an account.
type CashAccount11 struct {

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Identification *CashAccountIdentification1Choice `xml:"Id"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *FinancialInstitutionIdentification3Choice `xml:"AcctSvcr,omitempty"`
}

func (c *CashAccount11) AddIdentification() *CashAccountIdentification1Choice {
	c.Identification = new(CashAccountIdentification1Choice)
	return c.Identification
}

func (c *CashAccount11) AddAccountServicer() *FinancialInstitutionIdentification3Choice {
	c.AccountServicer = new(FinancialInstitutionIdentification3Choice)
	return c.AccountServicer
}
