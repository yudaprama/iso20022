package model

// Choice of formats for the specification of the party.
type Party23Choice struct {

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *Organisation16 `xml:"Org"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson23 `xml:"IndvPrsn"`
}

func (p *Party23Choice) AddOrganisation() *Organisation16 {
	p.Organisation = new(Organisation16)
	return p.Organisation
}

func (p *Party23Choice) AddIndividualPerson() *IndividualPerson23 {
	p.IndividualPerson = new(IndividualPerson23)
	return p.IndividualPerson
}
