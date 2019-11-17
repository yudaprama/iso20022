package model

// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
type Intermediary13 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account2 `xml:"Acct,omitempty"`

	// Non-enforcement of the right to all or part of a commission by the party entitled to the commission.
	WaivedTrailerCommissionIndicator *YesNoIndicator `xml:"WvdTrlrComssnInd,omitempty"`

	// Role or function performed by an intermediary in a given situation.
	Role *InvestmentFundRole3Code `xml:"Role,omitempty"`

	// Role or function performed by an intermediary in a given situation.
	ExtendedRole *Extended350Code `xml:"XtndedRole,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	NameAndAddress *NameAndAddress4 `xml:"NmAndAdr,omitempty"`
}

func (i *Intermediary13) AddIdentification() *PartyIdentification4Choice {
	i.Identification = new(PartyIdentification4Choice)
	return i.Identification
}

func (i *Intermediary13) AddAccount() *Account2 {
	i.Account = new(Account2)
	return i.Account
}

func (i *Intermediary13) SetWaivedTrailerCommissionIndicator(value string) {
	i.WaivedTrailerCommissionIndicator = (*YesNoIndicator)(&value)
}

func (i *Intermediary13) SetRole(value string) {
	i.Role = (*InvestmentFundRole3Code)(&value)
}

func (i *Intermediary13) SetExtendedRole(value string) {
	i.ExtendedRole = (*Extended350Code)(&value)
}

func (i *Intermediary13) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *Intermediary13) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *Intermediary13) AddNameAndAddress() *NameAndAddress4 {
	i.NameAndAddress = new(NameAndAddress4)
	return i.NameAndAddress
}
