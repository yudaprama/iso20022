package model

// Information which describes the organisation.
type Organisation12 struct {

	// Name by which a party is known and which is usually used to identify that party.
	FullLegalName *Max350Text `xml:"FullLglNm"`

	// Name used by a business for commercial purposes, although its registered legal name, used for contracts and other formal situations, may be another.
	TradingName *Max350Text `xml:"TradgNm,omitempty"`

	// Country in which the organisation has its business activity.
	CountryOfOperation *CountryCode `xml:"CtryOfOpr"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Is an operational address, for example, of a shared services center.
	OperationalAddress *PostalAddress6 `xml:"OprlAdr,omitempty"`

	// Is the address where the business activity is taking place.
	BusinessAddress *PostalAddress6 `xml:"BizAdr,omitempty"`

	// Is the address where the entity resides and is registered. More generically, it is the home address (Residential address).
	LegalAddress *PostalAddress6 `xml:"LglAdr"`

	// Address where invoices must be sent.
	BillingAddress *PostalAddress6 `xml:"BllgAdr,omitempty"`

	// Unique and unambiguous way of identifying an organisation.
	OrganisationIdentification *OrganisationIdentification8 `xml:"OrgId"`

	// Person in the customer's organisation who can be contacted by the account servicer in relation to the account(s) identified in this instruction.
	RepresentativeOfficer []*PartyIdentification40 `xml:"RprtvOffcr,omitempty"`

	// Person responsible of the treasury department within the customer’s organisation.
	TreasuryManager *PartyIdentification40 `xml:"TrsrMgr,omitempty"`

	// Person that has the mandate to delegate authority, to assign mandates to other individuals (mandate holders) to perform specific bank operations on the account.
	MainMandateHolder []*PartyIdentification40 `xml:"MainMndtHldr,omitempty"`

	// Person that may be the potential sender of a message related to the life cycle of the account.
	Sender []*PartyIdentification40 `xml:"Sndr,omitempty"`

	// Person that is officially and legally mandated to represent the organisation. Depending on legislation, the legal representative(s) might for instance be assigned by the Board, identified in the by-laws of the organisation, be publicly announced in the official journal of a country, etc.
	LegalRepresentative []*PartyIdentification40 `xml:"LglRprtv,omitempty"`
}

func (o *Organisation12) SetFullLegalName(value string) {
	o.FullLegalName = (*Max350Text)(&value)
}

func (o *Organisation12) SetTradingName(value string) {
	o.TradingName = (*Max350Text)(&value)
}

func (o *Organisation12) SetCountryOfOperation(value string) {
	o.CountryOfOperation = (*CountryCode)(&value)
}

func (o *Organisation12) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation12) AddOperationalAddress() *PostalAddress6 {
	o.OperationalAddress = new(PostalAddress6)
	return o.OperationalAddress
}

func (o *Organisation12) AddBusinessAddress() *PostalAddress6 {
	o.BusinessAddress = new(PostalAddress6)
	return o.BusinessAddress
}

func (o *Organisation12) AddLegalAddress() *PostalAddress6 {
	o.LegalAddress = new(PostalAddress6)
	return o.LegalAddress
}

func (o *Organisation12) AddBillingAddress() *PostalAddress6 {
	o.BillingAddress = new(PostalAddress6)
	return o.BillingAddress
}

func (o *Organisation12) AddOrganisationIdentification() *OrganisationIdentification8 {
	o.OrganisationIdentification = new(OrganisationIdentification8)
	return o.OrganisationIdentification
}

func (o *Organisation12) AddRepresentativeOfficer() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.RepresentativeOfficer = append(o.RepresentativeOfficer, newValue)
	return newValue
}

func (o *Organisation12) AddTreasuryManager() *PartyIdentification40 {
	o.TreasuryManager = new(PartyIdentification40)
	return o.TreasuryManager
}

func (o *Organisation12) AddMainMandateHolder() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.MainMandateHolder = append(o.MainMandateHolder, newValue)
	return newValue
}

func (o *Organisation12) AddSender() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.Sender = append(o.Sender, newValue)
	return newValue
}

func (o *Organisation12) AddLegalRepresentative() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.LegalRepresentative = append(o.LegalRepresentative, newValue)
	return newValue
}
