package model

// Reference information concerning the original transaction, to which the investigation message refers.
type UnderlyingPaymentTransaction2 struct {

	// Set of elements used to provide information on the original message.
	OriginalGroupInformation *UnderlyingGroupInformation1 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt"`
}

func (u *UnderlyingPaymentTransaction2) AddOriginalGroupInformation() *UnderlyingGroupInformation1 {
	u.OriginalGroupInformation = new(UnderlyingGroupInformation1)
	return u.OriginalGroupInformation
}

func (u *UnderlyingPaymentTransaction2) SetOriginalInstructionIdentification(value string) {
	u.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentTransaction2) SetOriginalEndToEndIdentification(value string) {
	u.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentTransaction2) SetOriginalTransactionIdentification(value string) {
	u.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentTransaction2) SetOriginalInterbankSettlementAmount(value, currency string) {
	u.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (u *UnderlyingPaymentTransaction2) SetOriginalInterbankSettlementDate(value string) {
	u.OriginalInterbankSettlementDate = (*ISODate)(&value)
}
