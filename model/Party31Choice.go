package model

// Choice of formats for the specification of the party.
type Party31Choice struct {

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *Organisation24 `xml:"Org"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson27 `xml:"IndvPrsn"`
}

func (p *Party31Choice) AddOrganisation() *Organisation24 {
	p.Organisation = new(Organisation24)
	return p.Organisation
}

func (p *Party31Choice) AddIndividualPerson() *IndividualPerson27 {
	p.IndividualPerson = new(IndividualPerson27)
	return p.IndividualPerson
}
