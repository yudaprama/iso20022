package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount14 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Identification of an individual person whom legally owns the account.
	IndividualOwnerIdentification *IndividualPersonIdentificationChoice `xml:"IndvOwnrId,omitempty"`

	// Identification of an organisation that legally owns the account.
	OrganisationOwnerIdentification *PartyIdentification2Choice `xml:"OrgOwnrId,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount14) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount14) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount14) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount14) AddIndividualOwnerIdentification() *IndividualPersonIdentificationChoice {
	i.IndividualOwnerIdentification = new(IndividualPersonIdentificationChoice)
	return i.IndividualOwnerIdentification
}

func (i *InvestmentAccount14) AddOrganisationOwnerIdentification() *PartyIdentification2Choice {
	i.OrganisationOwnerIdentification = new(PartyIdentification2Choice)
	return i.OrganisationOwnerIdentification
}

func (i *InvestmentAccount14) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}
