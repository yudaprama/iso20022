package model

// Provides the details of the identification information.
type IdentificationInformation2 struct {

	// Account owner that owes an amount of money or to whom an amount of money is due.
	Party *PartyIdentification43 `xml:"Pty,omitempty"`

	// Unambiguous identification of the account of a party.
	Account *AccountIdentification4Choice `xml:"Acct,omitempty"`

	// Financial institution servicing an account for a party.
	Agent *BranchAndFinancialInstitutionIdentification5 `xml:"Agt,omitempty"`
}

func (i *IdentificationInformation2) AddParty() *PartyIdentification43 {
	i.Party = new(PartyIdentification43)
	return i.Party
}

func (i *IdentificationInformation2) AddAccount() *AccountIdentification4Choice {
	i.Account = new(AccountIdentification4Choice)
	return i.Account
}

func (i *IdentificationInformation2) AddAgent() *BranchAndFinancialInstitutionIdentification5 {
	i.Agent = new(BranchAndFinancialInstitutionIdentification5)
	return i.Agent
}
