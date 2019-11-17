package model

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader71 struct {

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

	// Total amount of money moved between the instructing agent and the instructed agent in the reversal message.
	TotalReversedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlRvsdIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInstruction4 `xml:"SttlmInf"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the reversal message and not the party that received the original instruction that is being reversed.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader71) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader71) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader71) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader71) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader71) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader71) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader71) SetGroupReversal(value string) {
	g.GroupReversal = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader71) SetTotalReversedInterbankSettlementAmount(value, currency string) {
	g.TotalReversedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader71) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader71) AddSettlementInformation() *SettlementInstruction4 {
	g.SettlementInformation = new(SettlementInstruction4)
	return g.SettlementInformation
}

func (g *GroupHeader71) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructingAgent
}

func (g *GroupHeader71) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructedAgent
}
