package model

// Information used for identifying an account.
type CashAccount34 struct {

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Identification *AccountIdentificationAndName5 `xml:"Id"`

	// Institution that maintains the records where the account is held.
	AccountServicer *FinancialInstitutionIdentification7Choice `xml:"AcctSvcr,omitempty"`
}

func (c *CashAccount34) AddIdentification() *AccountIdentificationAndName5 {
	c.Identification = new(AccountIdentificationAndName5)
	return c.Identification
}

func (c *CashAccount34) AddAccountServicer() *FinancialInstitutionIdentification7Choice {
	c.AccountServicer = new(FinancialInstitutionIdentification7Choice)
	return c.AccountServicer
}
