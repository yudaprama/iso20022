package model

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader48 struct {

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

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Party that initiates the payment.
	//
	// Usage: This can either be the debtor or the party that initiates the credit transfer on behalf of the debtor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain for execution.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"FwdgAgt,omitempty"`
}

func (g *GroupHeader48) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader48) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader48) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader48) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader48) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader48) AddInitiatingParty() *PartyIdentification43 {
	g.InitiatingParty = new(PartyIdentification43)
	return g.InitiatingParty
}

func (g *GroupHeader48) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.ForwardingAgent
}
