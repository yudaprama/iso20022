package model

// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
type Intermediary7 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account2 `xml:"Acct,omitempty"`
}

func (i *Intermediary7) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary7) AddAccount() *Account2 {
	i.Account = new(Account2)
	return i.Account
}
