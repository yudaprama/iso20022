package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount42 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Party that legally owns the account.
	OwnerIdentification *PartyIdentification2Choice `xml:"OwnrId,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount42) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount42) AddOwnerIdentification() *PartyIdentification2Choice {
	i.OwnerIdentification = new(PartyIdentification2Choice)
	return i.OwnerIdentification
}

func (i *InvestmentAccount42) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}
