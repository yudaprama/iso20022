package model

// Choice of formats for the specification of the party.
type Party32Choice struct {

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *Organisation30 `xml:"Org"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson34 `xml:"IndvPrsn"`
}

func (p *Party32Choice) AddOrganisation() *Organisation30 {
	p.Organisation = new(Organisation30)
	return p.Organisation
}

func (p *Party32Choice) AddIndividualPerson() *IndividualPerson34 {
	p.IndividualPerson = new(IndividualPerson34)
	return p.IndividualPerson
}
