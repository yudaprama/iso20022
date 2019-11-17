package model

// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
type Organisation13 struct {

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

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress3 `xml:"PstlAdr"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Additional regulatory information about the investor that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`
}

func (o *Organisation13) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation13) AddIdentification() *PartyIdentification4Choice {
	o.Identification = new(PartyIdentification4Choice)
	return o.Identification
}

func (o *Organisation13) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation13) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation13) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation13) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation13) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *Organisation13) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation13) AddPostalAddress() *PostalAddress3 {
	newValue := new(PostalAddress3)
	o.PostalAddress = append(o.PostalAddress, newValue)
	return newValue
}

func (o *Organisation13) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	o.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return o.PrimaryCommunicationAddress
}

func (o *Organisation13) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	o.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return o.SecondaryCommunicationAddress
}

func (o *Organisation13) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	o.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return o.AdditionalRegulatoryInformation
}
