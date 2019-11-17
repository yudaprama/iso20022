package model

// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
type Organisation4 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id,omitempty"`

	// Purpose of the organisation, eg, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country of taxation of an organisation.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb,omitempty"`

	// Number assigned by a national registration authority to an entity.
	NationalRegistrationNumber *Max35Text `xml:"NtlRegnNb,omitempty"`

	// Postal address of a party.
	CorporateInvestorAddress *PostalAddress1 `xml:"CorpInvstrAdr"`
}

func (o *Organisation4) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation4) AddIdentification() *PartyIdentification4Choice {
	o.Identification = new(PartyIdentification4Choice)
	return o.Identification
}

func (o *Organisation4) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation4) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation4) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation4) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation4) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *Organisation4) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation4) AddCorporateInvestorAddress() *PostalAddress1 {
	o.CorporateInvestorAddress = new(PostalAddress1)
	return o.CorporateInvestorAddress
}
