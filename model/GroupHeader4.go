package model

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader4 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions in the message is requested.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money transferred between instructing agent and instructed agent.
	TotalInterbankSettlementAmount *CurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation1 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation5 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader4) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader4) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader4) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader4) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader4) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader4) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (g *GroupHeader4) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader4) AddSettlementInformation() *SettlementInformation1 {
	g.SettlementInformation = new(SettlementInformation1)
	return g.SettlementInformation
}

func (g *GroupHeader4) AddPaymentTypeInformation() *PaymentTypeInformation5 {
	g.PaymentTypeInformation = new(PaymentTypeInformation5)
	return g.PaymentTypeInformation
}

func (g *GroupHeader4) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructingAgent
}

func (g *GroupHeader4) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructedAgent
}
