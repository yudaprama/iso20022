package model

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader37 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader37) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader37) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader37) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructingAgent
}

func (g *GroupHeader37) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructedAgent
}
