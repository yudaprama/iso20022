package model

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation21 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Unique and unambiguous identifier for the organisation.
	Identification *PartyIdentification72Choice `xml:"Id,omitempty"`

	// Purpose of the organisation, for example,, charity.
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

func (o *Organisation21) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation21) AddIdentification() *PartyIdentification72Choice {
	o.Identification = new(PartyIdentification72Choice)
	return o.Identification
}

func (o *Organisation21) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation21) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation21) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation21) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation21) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *Organisation21) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation21) AddCorporateInvestorAddress() *PostalAddress1 {
	o.CorporateInvestorAddress = new(PostalAddress1)
	return o.CorporateInvestorAddress
}
