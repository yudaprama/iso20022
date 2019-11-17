package model

// Statement of the journal entries for all activities related to the registered currency control contract.
type ContractRegistrationStatement1 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the contract registration statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Party registering the currency control contract.
	ReportingParty *TradeParty2 `xml:"RptgPty"`

	// Agent which registers the currency control contract.
	RegistrationAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RegnAgt"`

	// Specifies the period for which the statement is provided.
	ReportingPeriod *ReportingPeriod1 `xml:"RptgPrd"`

	// Registered currency control contract.
	RegisteredContract *RegisteredContract5 `xml:"RegdCtrct"`

	// Journal of the transactions recorded under the registered currency control contract.
	TransactionJournal []*TransactionCertificate1 `xml:"TxJrnl,omitempty"`

	// Journal of the supporting documents recorded under the registered currency control contract.
	SupportingDocumentJournal []*SupportingDocument1 `xml:"SpprtgDocJrnl,omitempty"`

	// Journal of additional supporting documents recorded under the registered currency control contract.
	AdditionalSupportingDocumentJournal []*SupportingDocument1 `xml:"AddtlSpprtgDocJrnl,omitempty"`

	// Details on the currency control rule against which has been violated.
	RegulatoryRuleValidation []*GenericValidationRuleIdentification1 `xml:"RgltryRuleVldtn,omitempty"`

	// Total turn over amount recorded under the currency control contract for the amount of all
	TotalContractTurnoverSum *ActiveCurrencyAndAmount `xml:"TtlCtrctTrnvrSum"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *ContractRegistrationStatement1) SetStatementIdentification(value string) {
	c.StatementIdentification = (*Max35Text)(&value)
}

func (c *ContractRegistrationStatement1) AddReportingParty() *TradeParty2 {
	c.ReportingParty = new(TradeParty2)
	return c.ReportingParty
}

func (c *ContractRegistrationStatement1) AddRegistrationAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.RegistrationAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.RegistrationAgent
}

func (c *ContractRegistrationStatement1) AddReportingPeriod() *ReportingPeriod1 {
	c.ReportingPeriod = new(ReportingPeriod1)
	return c.ReportingPeriod
}

func (c *ContractRegistrationStatement1) AddRegisteredContract() *RegisteredContract5 {
	c.RegisteredContract = new(RegisteredContract5)
	return c.RegisteredContract
}

func (c *ContractRegistrationStatement1) AddTransactionJournal() *TransactionCertificate1 {
	newValue := new(TransactionCertificate1)
	c.TransactionJournal = append(c.TransactionJournal, newValue)
	return newValue
}

func (c *ContractRegistrationStatement1) AddSupportingDocumentJournal() *SupportingDocument1 {
	newValue := new(SupportingDocument1)
	c.SupportingDocumentJournal = append(c.SupportingDocumentJournal, newValue)
	return newValue
}

func (c *ContractRegistrationStatement1) AddAdditionalSupportingDocumentJournal() *SupportingDocument1 {
	newValue := new(SupportingDocument1)
	c.AdditionalSupportingDocumentJournal = append(c.AdditionalSupportingDocumentJournal, newValue)
	return newValue
}

func (c *ContractRegistrationStatement1) AddRegulatoryRuleValidation() *GenericValidationRuleIdentification1 {
	newValue := new(GenericValidationRuleIdentification1)
	c.RegulatoryRuleValidation = append(c.RegulatoryRuleValidation, newValue)
	return newValue
}

func (c *ContractRegistrationStatement1) SetTotalContractTurnoverSum(value, currency string) {
	c.TotalContractTurnoverSum = NewActiveCurrencyAndAmount(value, currency)
}

func (c *ContractRegistrationStatement1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
