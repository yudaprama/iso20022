package model

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader7 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the (group of) cancellation request(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the cancellation applies to a whole group of transactions or to individual transactions within the original group.
	GroupCancellation *GroupingIndicator `xml:"GrpCxl,omitempty"`

	// Party initiating the payment. In the payment context, this can either be the debtor (in a credit transfer), the creditor (in a direct debit), or the party that initiates the payment on behalf of the debtor or creditor.
	InitiatingParty *PartyIdentification8 `xml:"InitgPty,omitempty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"FwdgAgt,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader7) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader7) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader7) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader7) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader7) SetGroupCancellation(value string) {
	g.GroupCancellation = (*GroupingIndicator)(&value)
}

func (g *GroupHeader7) AddInitiatingParty() *PartyIdentification8 {
	g.InitiatingParty = new(PartyIdentification8)
	return g.InitiatingParty
}

func (g *GroupHeader7) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.ForwardingAgent
}

func (g *GroupHeader7) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.DebtorAgent
}

func (g *GroupHeader7) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.CreditorAgent
}

func (g *GroupHeader7) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructingAgent
}

func (g *GroupHeader7) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructedAgent
}
