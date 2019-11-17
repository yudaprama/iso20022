package model

// Choice of formats for the specification of the party.
type Party24Choice struct {

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *Organisation17 `xml:"Org"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson24 `xml:"IndvPrsn"`
}

func (p *Party24Choice) AddOrganisation() *Organisation17 {
	p.Organisation = new(Organisation17)
	return p.Organisation
}

func (p *Party24Choice) AddIndividualPerson() *IndividualPerson24 {
	p.IndividualPerson = new(IndividualPerson24)
	return p.IndividualPerson
}
