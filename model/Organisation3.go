package model

// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
type Organisation3 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id,omitempty"`

	// Purpose of the organisation, eg, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country of taxation of an individual person or an organisation.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb,omitempty"`

	// Number assigned by a national registration authority to an entity.
	NationalRegistrationNumber *Max35Text `xml:"NtlRegnNb,omitempty"`

	// Address information to be either inserted, updated or deleted.
	ModifiedPostalAddress []*ModificationScope1 `xml:"ModfdPstlAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`
}

func (o *Organisation3) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation3) AddIdentification() *PartyIdentification4Choice {
	o.Identification = new(PartyIdentification4Choice)
	return o.Identification
}

func (o *Organisation3) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation3) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation3) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation3) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation3) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *Organisation3) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation3) AddModifiedPostalAddress() *ModificationScope1 {
	newValue := new(ModificationScope1)
	o.ModifiedPostalAddress = append(o.ModifiedPostalAddress, newValue)
	return newValue
}

func (o *Organisation3) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	o.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return o.PrimaryCommunicationAddress
}

func (o *Organisation3) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	o.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return o.SecondaryCommunicationAddress
}
