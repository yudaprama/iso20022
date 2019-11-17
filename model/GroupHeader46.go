package model

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader46 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the status report was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party initiating the creditor payment activation request. This can either be the creditor himself or the party that initiates the request on behalf of the creditor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`
}

func (g *GroupHeader46) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader46) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader46) AddInitiatingParty() *PartyIdentification43 {
	g.InitiatingParty = new(PartyIdentification43)
	return g.InitiatingParty
}

func (g *GroupHeader46) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.DebtorAgent
}

func (g *GroupHeader46) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.CreditorAgent
}
