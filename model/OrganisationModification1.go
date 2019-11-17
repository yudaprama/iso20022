package model

// Information which describes the organisation.
type OrganisationModification1 struct {

	// Name by which a party is known and which is usually used to identify that party.
	FullLegalName *FullLegalNameModification1 `xml:"FullLglNm"`

	// Name used by a business for commercial purposes, although its registered legal name, used for contracts and other formal situations, may be another.
	TradingName *TradingNameModification1 `xml:"TradgNm,omitempty"`

	// Country in which the organisation has its business activity.
	CountryOfOperation *CountryCode `xml:"CtryOfOpr"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Is an operational address, for example, of a shared services center.
	OperationalAddress *AddressModification1 `xml:"OprlAdr,omitempty"`

	// Is the address where the business activity is taking place.
	BusinessAddress *AddressModification1 `xml:"BizAdr,omitempty"`

	// Is the address where the entity resides and is registered. More generically, it is the home address (Residential address).
	LegalAddress *AddressModification1 `xml:"LglAdr"`

	// Address where invoices must be sent.
	BillingAddress *AddressModification1 `xml:"BllgAdr,omitempty"`

	// Unique and unambiguous way of identifying an organisation.
	OrganisationIdentification *OrganisationIdentification8 `xml:"OrgId"`

	// Person in the customer's organisation who can be contacted by the account servicer in relation to the account(s) identified in this instruction.
	RepresentativeOfficer []*PartyModification1 `xml:"RprtvOffcr,omitempty"`

	// Person responsible of the treasury department within the customer’s organisation.
	TreasuryManager *PartyModification1 `xml:"TrsrMgr,omitempty"`

	// Person that has the mandate to delegate authority, to assign mandates to other individuals (mandate holders) to perform specific bank operations on the account.
	MainMandateHolder []*PartyModification1 `xml:"MainMndtHldr,omitempty"`

	// Person that may be the potential sender of a message related to the life cycle of the account.
	Sender []*PartyModification1 `xml:"Sndr,omitempty"`

	// Person that is officially and legally mandated to represent the organisation. Depending on legislation, the legal representative(s) might for instance be assigned by the Board, identified in the by-laws of the organisation, be publicly announced in the official journal of a country, etc.
	LegalRepresentative []*PartyModification1 `xml:"LglRprtv,omitempty"`
}

func (o *OrganisationModification1) AddFullLegalName() *FullLegalNameModification1 {
	o.FullLegalName = new(FullLegalNameModification1)
	return o.FullLegalName
}

func (o *OrganisationModification1) AddTradingName() *TradingNameModification1 {
	o.TradingName = new(TradingNameModification1)
	return o.TradingName
}

func (o *OrganisationModification1) SetCountryOfOperation(value string) {
	o.CountryOfOperation = (*CountryCode)(&value)
}

func (o *OrganisationModification1) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *OrganisationModification1) AddOperationalAddress() *AddressModification1 {
	o.OperationalAddress = new(AddressModification1)
	return o.OperationalAddress
}

func (o *OrganisationModification1) AddBusinessAddress() *AddressModification1 {
	o.BusinessAddress = new(AddressModification1)
	return o.BusinessAddress
}

func (o *OrganisationModification1) AddLegalAddress() *AddressModification1 {
	o.LegalAddress = new(AddressModification1)
	return o.LegalAddress
}

func (o *OrganisationModification1) AddBillingAddress() *AddressModification1 {
	o.BillingAddress = new(AddressModification1)
	return o.BillingAddress
}

func (o *OrganisationModification1) AddOrganisationIdentification() *OrganisationIdentification8 {
	o.OrganisationIdentification = new(OrganisationIdentification8)
	return o.OrganisationIdentification
}

func (o *OrganisationModification1) AddRepresentativeOfficer() *PartyModification1 {
	newValue := new(PartyModification1)
	o.RepresentativeOfficer = append(o.RepresentativeOfficer, newValue)
	return newValue
}

func (o *OrganisationModification1) AddTreasuryManager() *PartyModification1 {
	o.TreasuryManager = new(PartyModification1)
	return o.TreasuryManager
}

func (o *OrganisationModification1) AddMainMandateHolder() *PartyModification1 {
	newValue := new(PartyModification1)
	o.MainMandateHolder = append(o.MainMandateHolder, newValue)
	return newValue
}

func (o *OrganisationModification1) AddSender() *PartyModification1 {
	newValue := new(PartyModification1)
	o.Sender = append(o.Sender, newValue)
	return newValue
}

func (o *OrganisationModification1) AddLegalRepresentative() *PartyModification1 {
	newValue := new(PartyModification1)
	o.LegalRepresentative = append(o.LegalRepresentative, newValue)
	return newValue
}
