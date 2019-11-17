package model

// Choice of individual or organisation details.
type OwnerIdentification1Choice struct {

	// Identification of an individual person.
	IndividualOwnerIdentification *IndividualPersonIdentificationChoice `xml:"IndvOwnrId"`

	// Identification of an organisation.
	OrganisationOwnerIdentification *PartyIdentification5Choice `xml:"OrgOwnrId"`
}

func (o *OwnerIdentification1Choice) AddIndividualOwnerIdentification() *IndividualPersonIdentificationChoice {
	o.IndividualOwnerIdentification = new(IndividualPersonIdentificationChoice)
	return o.IndividualOwnerIdentification
}

func (o *OwnerIdentification1Choice) AddOrganisationOwnerIdentification() *PartyIdentification5Choice {
	o.OrganisationOwnerIdentification = new(PartyIdentification5Choice)
	return o.OrganisationOwnerIdentification
}
