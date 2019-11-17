package model

// Choice of formats for the specification of the party.
type Party30Choice struct {

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *Organisation22 `xml:"Org"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson28 `xml:"IndvPrsn"`
}

func (p *Party30Choice) AddOrganisation() *Organisation22 {
	p.Organisation = new(Organisation22)
	return p.Organisation
}

func (p *Party30Choice) AddIndividualPerson() *IndividualPerson28 {
	p.IndividualPerson = new(IndividualPerson28)
	return p.IndividualPerson
}
