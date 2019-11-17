package model

// Choice of formats for the specification of the party.
type Party33Choice struct {

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *Organisation29 `xml:"Org"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson33 `xml:"IndvPrsn"`
}

func (p *Party33Choice) AddOrganisation() *Organisation29 {
	p.Organisation = new(Organisation29)
	return p.Organisation
}

func (p *Party33Choice) AddIndividualPerson() *IndividualPerson33 {
	p.IndividualPerson = new(IndividualPerson33)
	return p.IndividualPerson
}
