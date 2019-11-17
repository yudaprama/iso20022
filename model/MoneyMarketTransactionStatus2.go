package model

// Provides the details of each individual secured market transaction.
type MoneyMarketTransactionStatus2 struct {

	// Unique transaction identifier will be created at the time a transaction is first executed, shared with all registered entities and counterparties involved in the transaction, and used to track that particular transaction during its lifetime.
	UniqueTransactionIdentifier *Max105Text `xml:"UnqTxIdr,omitempty"`

	// Internal unique transaction identifier used by the reporting agent for each transaction.
	ProprietaryTransactionIdentification *Max105Text `xml:"PrtryTxId"`

	// Unique and unambiguous legal entity identification of  the branch of the reporting agent in which the transaction has been booked.
	//
	// Usage: This field must only be provided if the transaction has been conducted and booked by a branch of the reporting agent and only if this branch has its own LEI that the reporting agent can clearly identify.
	// Where the transaction has been booked by the head office or the reporting agent cannot be identified by a unique branch-specific LEI, the reporting agent must provide the LEI of the head office.
	BranchIdentification *LEIIdentifier `xml:"BrnchId,omitempty"`

	// Defines status of the reported transaction.
	Status *StatisticalReportingStatus2Code `xml:"Sts"`

	// Provides the details of the rule which could not be validated.
	ValidationRule []*GenericValidationRuleIdentification1 `xml:"VldtnRule,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MoneyMarketTransactionStatus2) SetUniqueTransactionIdentifier(value string) {
	m.UniqueTransactionIdentifier = (*Max105Text)(&value)
}

func (m *MoneyMarketTransactionStatus2) SetProprietaryTransactionIdentification(value string) {
	m.ProprietaryTransactionIdentification = (*Max105Text)(&value)
}

func (m *MoneyMarketTransactionStatus2) SetBranchIdentification(value string) {
	m.BranchIdentification = (*LEIIdentifier)(&value)
}

func (m *MoneyMarketTransactionStatus2) SetStatus(value string) {
	m.Status = (*StatisticalReportingStatus2Code)(&value)
}

func (m *MoneyMarketTransactionStatus2) AddValidationRule() *GenericValidationRuleIdentification1 {
	newValue := new(GenericValidationRuleIdentification1)
	m.ValidationRule = append(m.ValidationRule, newValue)
	return newValue
}

func (m *MoneyMarketTransactionStatus2) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
