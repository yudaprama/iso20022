package model

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation24 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Name of the organisation in short form.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Unique and unambiguous identifier for the organisation.
	Identification *PartyIdentification72Choice `xml:"Id,omitempty"`

	// Identification of the organisation with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Purpose of the organisation, for example, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress21 `xml:"PstlAdr"`

	// Type of organisation.
	TypeOfOrganisation *OrganisationType1Choice `xml:"TpOfOrg,omitempty"`
}

func (o *Organisation24) SetName(value string) {
	o.Name = (*Max350Text)(&value)
}

func (o *Organisation24) SetShortName(value string) {
	o.ShortName = (*Max35Text)(&value)
}

func (o *Organisation24) AddIdentification() *PartyIdentification72Choice {
	o.Identification = new(PartyIdentification72Choice)
	return o.Identification
}

func (o *Organisation24) SetLegalEntityIdentifier(value string) {
	o.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (o *Organisation24) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation24) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation24) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation24) AddPostalAddress() *PostalAddress21 {
	newValue := new(PostalAddress21)
	o.PostalAddress = append(o.PostalAddress, newValue)
	return newValue
}

func (o *Organisation24) AddTypeOfOrganisation() *OrganisationType1Choice {
	o.TypeOfOrganisation = new(OrganisationType1Choice)
	return o.TypeOfOrganisation
}
