package model

// Choice of formats for the specification of the party.
type Party15Choice struct {

	// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
	Organisation *Organisation15 `xml:"Org"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson21 `xml:"IndvPrsn"`
}

func (p *Party15Choice) AddOrganisation() *Organisation15 {
	p.Organisation = new(Organisation15)
	return p.Organisation
}

func (p *Party15Choice) AddIndividualPerson() *IndividualPerson21 {
	p.IndividualPerson = new(IndividualPerson21)
	return p.IndividualPerson
}
