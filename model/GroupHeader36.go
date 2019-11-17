package model

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader36 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party that initiates the status message.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty,omitempty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"FwdgAgt,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`
}

func (g *GroupHeader36) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader36) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader36) AddInitiatingParty() *PartyIdentification32 {
	g.InitiatingParty = new(PartyIdentification32)
	return g.InitiatingParty
}

func (g *GroupHeader36) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.ForwardingAgent
}

func (g *GroupHeader36) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.DebtorAgent
}

func (g *GroupHeader36) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.CreditorAgent
}
