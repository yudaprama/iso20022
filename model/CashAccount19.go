package model

// Provides information about the cash account.
type CashAccount19 struct {

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Identification of the cash account or the securities account from which the cash account is derived.
	AccountIdentification *AccountIdentification2Choice `xml:"AcctId"`
}

func (c *CashAccount19) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashAccount19) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	c.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return c.AccountOwnerIdentification
}

func (c *CashAccount19) AddAccountIdentification() *AccountIdentification2Choice {
	c.AccountIdentification = new(AccountIdentification2Choice)
	return c.AccountIdentification
}
