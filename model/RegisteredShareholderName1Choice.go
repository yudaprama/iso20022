package model

// Choice of individual or organisation name and address.
type RegisteredShareholderName1Choice struct {

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson29 `xml:"IndvPrsn"`

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *Organisation23 `xml:"Org"`
}

func (r *RegisteredShareholderName1Choice) AddIndividualPerson() *IndividualPerson29 {
	r.IndividualPerson = new(IndividualPerson29)
	return r.IndividualPerson
}

func (r *RegisteredShareholderName1Choice) AddOrganisation() *Organisation23 {
	r.Organisation = new(Organisation23)
	return r.Organisation
}
