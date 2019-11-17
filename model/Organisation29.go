package model

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation29 struct {

	// Name by which the organisation is known and which is usually used to identify that  organisation.
	Name *Max350Text `xml:"Nm,omitempty"`

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

	// Information related to an address to be inserted, updated or deleted.
	ModifiedPostalAddress []*ModificationScope34 `xml:"ModfdPstlAdr,omitempty"`

	// Type of organisation.
	TypeOfOrganisation *OrganisationType1Choice `xml:"TpOfOrg,omitempty"`

	// Place of listing for shares in the organisation.
	PlaceOfListing []*MICIdentifier `xml:"PlcOfListg,omitempty"`
}

func (o *Organisation29) SetName(value string) {
	o.Name = (*Max350Text)(&value)
}

func (o *Organisation29) SetShortName(value string) {
	o.ShortName = (*Max35Text)(&value)
}

func (o *Organisation29) AddIdentification() *PartyIdentification72Choice {
	o.Identification = new(PartyIdentification72Choice)
	return o.Identification
}

func (o *Organisation29) SetLegalEntityIdentifier(value string) {
	o.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (o *Organisation29) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation29) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation29) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation29) AddModifiedPostalAddress() *ModificationScope34 {
	newValue := new(ModificationScope34)
	o.ModifiedPostalAddress = append(o.ModifiedPostalAddress, newValue)
	return newValue
}

func (o *Organisation29) AddTypeOfOrganisation() *OrganisationType1Choice {
	o.TypeOfOrganisation = new(OrganisationType1Choice)
	return o.TypeOfOrganisation
}

func (o *Organisation29) AddPlaceOfListing(value string) {
	o.PlaceOfListing = append(o.PlaceOfListing, (*MICIdentifier)(&value))
}
