package model

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader31 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the instructed party, to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check the authority of the initiating party.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side. The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Party that initiates the mandate message.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty,omitempty"`

	// Agent that instructs the next party in the chain to carry out an instruction.
	//
	// Usage Rule:
	// In case of amendment and cancellation request messages, the instructing agent is the party sending the amendment and cancellation request message and not the party that sent the original mandate initiation request message.
	// In case of acceptance report message, the instructing agent is the party sending the acceptance report message and not the party that sent the original mandate request message.
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out an instruction.
	//
	// Usage Rule:
	// In case of amendment and cancellation request messages, the instructed agent is the party receiving the amendment and cancellation request message and not the party that received the original mandate initiation request message.
	// In case of acceptance report message, the instructed agent is the party receiving the acceptance report message and not the party that received the original mandate request message.
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader31) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader31) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader31) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader31) AddInitiatingParty() *PartyIdentification32 {
	g.InitiatingParty = new(PartyIdentification32)
	return g.InitiatingParty
}

func (g *GroupHeader31) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructingAgent
}

func (g *GroupHeader31) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructedAgent
}
