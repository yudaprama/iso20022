package model

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation15 struct {

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
	ModifiedPostalAddress []*ModificationScope1 `xml:"ModfdPstlAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Additional regulatory information about the investor that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`
}

func (o *Organisation15) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation15) AddIdentification() *PartyIdentification4Choice {
	o.Identification = new(PartyIdentification4Choice)
	return o.Identification
}

func (o *Organisation15) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation15) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation15) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation15) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation15) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *Organisation15) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation15) AddModifiedPostalAddress() *ModificationScope1 {
	newValue := new(ModificationScope1)
	o.ModifiedPostalAddress = append(o.ModifiedPostalAddress, newValue)
	return newValue
}

func (o *Organisation15) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	o.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return o.PrimaryCommunicationAddress
}

func (o *Organisation15) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	o.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return o.SecondaryCommunicationAddress
}

func (o *Organisation15) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	o.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return o.AdditionalRegulatoryInformation
}
