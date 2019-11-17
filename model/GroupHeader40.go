package model

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader40 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the reversal applies to the whole group of transactions or to individual transactions within the original group.
	GroupReversal *TrueFalseIndicator `xml:"GrpRvsl,omitempty"`

	// Party that initiates the reversal message.
	// Usage: This can be either the creditor or a party that initiates the reversal of the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty,omitempty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"FwdgAgt,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`
}

func (g *GroupHeader40) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader40) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader40) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader40) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader40) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader40) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader40) SetGroupReversal(value string) {
	g.GroupReversal = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader40) AddInitiatingParty() *PartyIdentification32 {
	g.InitiatingParty = new(PartyIdentification32)
	return g.InitiatingParty
}

func (g *GroupHeader40) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.ForwardingAgent
}

func (g *GroupHeader40) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.DebtorAgent
}

func (g *GroupHeader40) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.CreditorAgent
}
