package model

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader8 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key that allows to check if the initiating party is allowed to initiate transactions from the account specified in the initiation.
	//
	// Usage: the content is not of a technical nature, but reflects the organisational structure at the initiating side. The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a different party than the initiating party.
	Authorisation []*Max128Text `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions in the message is requested.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the reversal applies to the whole group of transactions or to individual transactions within the original group.
	GroupReversal *TrueFalseIndicator `xml:"GrpRvsl,omitempty"`

	// Party initiating the payment. In the payment context, this can either be the debtor (in a credit transfer), the creditor (in a direct debit), or the party that initiates the payment on behalf of the debtor or creditor.
	InitiatingParty *PartyIdentification8 `xml:"InitgPty,omitempty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"FwdgAgt,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt,omitempty"`
}

func (g *GroupHeader8) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader8) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader8) AddAuthorisation(value string) {
	g.Authorisation = append(g.Authorisation, (*Max128Text)(&value))
}

func (g *GroupHeader8) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader8) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader8) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader8) SetGroupReversal(value string) {
	g.GroupReversal = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader8) AddInitiatingParty() *PartyIdentification8 {
	g.InitiatingParty = new(PartyIdentification8)
	return g.InitiatingParty
}

func (g *GroupHeader8) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.ForwardingAgent
}

func (g *GroupHeader8) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.DebtorAgent
}

func (g *GroupHeader8) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.CreditorAgent
}
