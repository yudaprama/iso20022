package model

// Information which describes the organisation.
type Organisation7 struct {

	// Name by which a party is known and which is usually used to identify that party.
	FullLegalName *Max350Text `xml:"FullLglNm"`

	// Name used by a business for commercial purposes, although its registered legal name, used for contracts and other formal situations, may be another.
	TradingName *Max350Text `xml:"TradgNm,omitempty"`

	// Country in which the organisation has its business activity.
	CountryOfOperation *CountryCode `xml:"CtryOfOpr"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt"`

	// Is an operational address, for example, of a shared services center.
	OperationalAddress *PostalAddress6 `xml:"OprlAdr,omitempty"`

	// Is the address where the business activity is taking place.
	BusinessAddress *PostalAddress6 `xml:"BizAdr,omitempty"`

	// Is the address where the entity resides and is registered. More generically, it is the home address (Residential address).
	LegalAddress *PostalAddress6 `xml:"LglAdr"`

	// Unique and unambiguous way of identifying an organisation.
	OrganisationIdentification *OrganisationIdentification6 `xml:"OrgId"`

	// Person in the customer's organisation who can be contacted by the account servicer.
	RepresentativeOfficer []*PartyIdentification40 `xml:"RprtvOffcr,omitempty"`

	// Identification of the person responsible of the treasury department within an organisation.
	TreasuryManager *PartyIdentification40 `xml:"TrsrMgr,omitempty"`

	// Is the main mandate holder that will delegate some authority to other individuals (mandate holders) to perform some specific bank operations on the account.
	MainMandateHolder []*PartyIdentification40 `xml:"MainMndtHldr,omitempty"`

	// Potential sender of a message related to the life cyle of an account.
	Sender []*PartyIdentification40 `xml:"Sndr,omitempty"`
}

func (o *Organisation7) SetFullLegalName(value string) {
	o.FullLegalName = (*Max350Text)(&value)
}

func (o *Organisation7) SetTradingName(value string) {
	o.TradingName = (*Max350Text)(&value)
}

func (o *Organisation7) SetCountryOfOperation(value string) {
	o.CountryOfOperation = (*CountryCode)(&value)
}

func (o *Organisation7) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation7) AddOperationalAddress() *PostalAddress6 {
	o.OperationalAddress = new(PostalAddress6)
	return o.OperationalAddress
}

func (o *Organisation7) AddBusinessAddress() *PostalAddress6 {
	o.BusinessAddress = new(PostalAddress6)
	return o.BusinessAddress
}

func (o *Organisation7) AddLegalAddress() *PostalAddress6 {
	o.LegalAddress = new(PostalAddress6)
	return o.LegalAddress
}

func (o *Organisation7) AddOrganisationIdentification() *OrganisationIdentification6 {
	o.OrganisationIdentification = new(OrganisationIdentification6)
	return o.OrganisationIdentification
}

func (o *Organisation7) AddRepresentativeOfficer() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.RepresentativeOfficer = append(o.RepresentativeOfficer, newValue)
	return newValue
}

func (o *Organisation7) AddTreasuryManager() *PartyIdentification40 {
	o.TreasuryManager = new(PartyIdentification40)
	return o.TreasuryManager
}

func (o *Organisation7) AddMainMandateHolder() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.MainMandateHolder = append(o.MainMandateHolder, newValue)
	return newValue
}

func (o *Organisation7) AddSender() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.Sender = append(o.Sender, newValue)
	return newValue
}
