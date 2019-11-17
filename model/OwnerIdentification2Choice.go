package model

// Choice of individual or organisation details.
type OwnerIdentification2Choice struct {

	// Identification of the individual person.
	IndividualOwnerIdentification *IndividualPersonIdentification2Choice `xml:"IndvOwnrId"`

	// Identification of an organisation.
	OrganisationOwnerIdentification *PartyIdentification95 `xml:"OrgOwnrId"`
}

func (o *OwnerIdentification2Choice) AddIndividualOwnerIdentification() *IndividualPersonIdentification2Choice {
	o.IndividualOwnerIdentification = new(IndividualPersonIdentification2Choice)
	return o.IndividualOwnerIdentification
}

func (o *OwnerIdentification2Choice) AddOrganisationOwnerIdentification() *PartyIdentification95 {
	o.OrganisationOwnerIdentification = new(PartyIdentification95)
	return o.OrganisationOwnerIdentification
}
