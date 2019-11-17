package model

// Choice of formats for the specification of the party.
type Party14Choice struct {

	// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
	Organisation *Organisation13 `xml:"Org"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson20 `xml:"IndvPrsn"`
}

func (p *Party14Choice) AddOrganisation() *Organisation13 {
	p.Organisation = new(Organisation13)
	return p.Organisation
}

func (p *Party14Choice) AddIndividualPerson() *IndividualPerson20 {
	p.IndividualPerson = new(IndividualPerson20)
	return p.IndividualPerson
}
