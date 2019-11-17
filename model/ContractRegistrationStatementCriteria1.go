package model

// Provides the list of criteria to be returned in the contract registration statement.
type ContractRegistrationStatementCriteria1 struct {

	// Indicates whether the journal of the transactions recorded under the registered currency control contract must be returned or not.
	TransactionJournal *TrueFalseIndicator `xml:"TxJrnl,omitempty"`

	// Indicates whether the journal of the supporting documents recorded under the registered currency control contract must be returned or not.
	SupportingDocumentJournal *TrueFalseIndicator `xml:"SpprtgDocJrnl,omitempty"`

	// Indicates whether the journal of additional supporting documents recorded under the registered currency control contract must be returned or not.
	AdditionalSupportingDocumentJournal *TrueFalseIndicator `xml:"AddtlSpprtgDocJrnl,omitempty"`

	// Indicates whether the details on the currency control rule against which has been violated must be returned or not.
	RegulatoryRuleValidation *TrueFalseIndicator `xml:"RgltryRuleVldtn,omitempty"`
}

func (c *ContractRegistrationStatementCriteria1) SetTransactionJournal(value string) {
	c.TransactionJournal = (*TrueFalseIndicator)(&value)
}

func (c *ContractRegistrationStatementCriteria1) SetSupportingDocumentJournal(value string) {
	c.SupportingDocumentJournal = (*TrueFalseIndicator)(&value)
}

func (c *ContractRegistrationStatementCriteria1) SetAdditionalSupportingDocumentJournal(value string) {
	c.AdditionalSupportingDocumentJournal = (*TrueFalseIndicator)(&value)
}

func (c *ContractRegistrationStatementCriteria1) SetRegulatoryRuleValidation(value string) {
	c.RegulatoryRuleValidation = (*TrueFalseIndicator)(&value)
}
