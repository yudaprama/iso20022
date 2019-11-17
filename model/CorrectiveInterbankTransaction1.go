package model

// Set of elements used to provide information on the corrective interbank transaction, to which the resolution message refers.
type CorrectiveInterbankTransaction1 struct {

	// Set of elements used to provide corrective information for the group header of the message under investigation.
	GroupHeader *CorrectiveGroupInformation1 `xml:"GrpHdr,omitempty"`

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level.
	// Usage: The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt"`
}

func (c *CorrectiveInterbankTransaction1) AddGroupHeader() *CorrectiveGroupInformation1 {
	c.GroupHeader = new(CorrectiveGroupInformation1)
	return c.GroupHeader
}

func (c *CorrectiveInterbankTransaction1) SetInstructionIdentification(value string) {
	c.InstructionIdentification = (*Max35Text)(&value)
}

func (c *CorrectiveInterbankTransaction1) SetEndToEndIdentification(value string) {
	c.EndToEndIdentification = (*Max35Text)(&value)
}

func (c *CorrectiveInterbankTransaction1) SetTransactionIdentification(value string) {
	c.TransactionIdentification = (*Max35Text)(&value)
}

func (c *CorrectiveInterbankTransaction1) SetInterbankSettlementAmount(value, currency string) {
	c.InterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CorrectiveInterbankTransaction1) SetInterbankSettlementDate(value string) {
	c.InterbankSettlementDate = (*ISODate)(&value)
}
