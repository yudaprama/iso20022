package model

// Choice between an organisation and an individual person.
type AccountOwner2Choice struct {

	// Identification of the individual person that legally owns the account.
	IndividualOwnerIdentification *IndividualPersonIdentification3Choice `xml:"IndvOwnrId"`

	// Identification of the organisation that legally owns the account.
	OrganisationOwnerIdentification *PartyIdentification96 `xml:"OrgOwnrId"`
}

func (a *AccountOwner2Choice) AddIndividualOwnerIdentification() *IndividualPersonIdentification3Choice {
	a.IndividualOwnerIdentification = new(IndividualPersonIdentification3Choice)
	return a.IndividualOwnerIdentification
}

func (a *AccountOwner2Choice) AddOrganisationOwnerIdentification() *PartyIdentification96 {
	a.OrganisationOwnerIdentification = new(PartyIdentification96)
	return a.OrganisationOwnerIdentification
}
