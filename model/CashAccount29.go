package model

// Information used for identifying an account.
type CashAccount29 struct {

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Identification *CashAccountIdentification1Choice `xml:"Id"`

	// Institution that maintains the records where the account is held.
	AccountServicer *FinancialInstitutionIdentification3Choice `xml:"AcctSvcr,omitempty"`
}

func (c *CashAccount29) AddIdentification() *CashAccountIdentification1Choice {
	c.Identification = new(CashAccountIdentification1Choice)
	return c.Identification
}

func (c *CashAccount29) AddAccountServicer() *FinancialInstitutionIdentification3Choice {
	c.AccountServicer = new(FinancialInstitutionIdentification3Choice)
	return c.AccountServicer
}
