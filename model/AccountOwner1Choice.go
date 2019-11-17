package model

// Choice between an organisation and an individual person.
type AccountOwner1Choice struct {

	// Identification of the individual person that legally owns the account.
	IndividualOwnerIdentification *IndividualPersonIdentification1Choice `xml:"IndvOwnrId"`

	// Identification of the organisation that legally owns the account.
	OrganisationOwnerIdentification *PartyIdentification96 `xml:"OrgOwnrId"`
}

func (a *AccountOwner1Choice) AddIndividualOwnerIdentification() *IndividualPersonIdentification1Choice {
	a.IndividualOwnerIdentification = new(IndividualPersonIdentification1Choice)
	return a.IndividualOwnerIdentification
}

func (a *AccountOwner1Choice) AddOrganisationOwnerIdentification() *PartyIdentification96 {
	a.OrganisationOwnerIdentification = new(PartyIdentification96)
	return a.OrganisationOwnerIdentification
}
