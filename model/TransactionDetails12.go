package model

// Identifies the details of the transaction.
type TransactionDetails12 struct {

	// Reference to the message advised to be cancelled by the account servicer
	Reference *References3Choice `xml:"Ref"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`
}

func (t *TransactionDetails12) AddReference() *References3Choice {
	t.Reference = new(References3Choice)
	return t.Reference
}

func (t *TransactionDetails12) AddAccountOwner() *PartyIdentification13Choice {
	t.AccountOwner = new(PartyIdentification13Choice)
	return t.AccountOwner
}

func (t *TransactionDetails12) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}
