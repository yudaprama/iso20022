package model

// Identification of a party and its role.
type Intermediary36 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution.
	Identification *PartyIdentification72Choice `xml:"Id"`

	// Identification of the organisation with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account20 `xml:"Acct,omitempty"`

	// Non-enforcement of the right to all or part of a commission by the party entitled to the commission.
	WaivedTrailerCommissionIndicator *YesNoIndicator `xml:"WvdTrlrComssnInd,omitempty"`

	// Role or function performed by the intermediary.
	Role *PartyRole2Choice `xml:"Role,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress []*CommunicationAddress6 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress []*CommunicationAddress6 `xml:"ScndryComAdr,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	NameAndAddress *NameAndAddress4 `xml:"NmAndAdr,omitempty"`
}

func (i *Intermediary36) AddIdentification() *PartyIdentification72Choice {
	i.Identification = new(PartyIdentification72Choice)
	return i.Identification
}

func (i *Intermediary36) SetLegalEntityIdentifier(value string) {
	i.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (i *Intermediary36) AddAccount() *Account20 {
	i.Account = new(Account20)
	return i.Account
}

func (i *Intermediary36) SetWaivedTrailerCommissionIndicator(value string) {
	i.WaivedTrailerCommissionIndicator = (*YesNoIndicator)(&value)
}

func (i *Intermediary36) AddRole() *PartyRole2Choice {
	i.Role = new(PartyRole2Choice)
	return i.Role
}

func (i *Intermediary36) AddPrimaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.PrimaryCommunicationAddress = append(i.PrimaryCommunicationAddress, newValue)
	return newValue
}

func (i *Intermediary36) AddSecondaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.SecondaryCommunicationAddress = append(i.SecondaryCommunicationAddress, newValue)
	return newValue
}

func (i *Intermediary36) AddNameAndAddress() *NameAndAddress4 {
	i.NameAndAddress = new(NameAndAddress4)
	return i.NameAndAddress
}
