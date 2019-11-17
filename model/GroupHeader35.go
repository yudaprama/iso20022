package model

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader35 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent.
	TotalInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation13 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation23 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader35) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader35) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader35) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader35) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader35) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader35) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader35) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader35) AddSettlementInformation() *SettlementInformation13 {
	g.SettlementInformation = new(SettlementInformation13)
	return g.SettlementInformation
}

func (g *GroupHeader35) AddPaymentTypeInformation() *PaymentTypeInformation23 {
	g.PaymentTypeInformation = new(PaymentTypeInformation23)
	return g.PaymentTypeInformation
}

func (g *GroupHeader35) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructingAgent
}

func (g *GroupHeader35) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructedAgent
}
