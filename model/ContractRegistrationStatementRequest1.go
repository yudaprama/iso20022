package model

// Details on the statement requested elements.
type ContractRegistrationStatementRequest1 struct {

	// Unique and unambiguous identification of the contract registration statement request.
	StatementRequestIdentification *Max35Text `xml:"StmtReqId"`

	// Specifies the period for which the statement is requested.
	ReportingPeriod *ReportingPeriod1 `xml:"RptgPrd"`

	// Party registering the currency control contract.
	ReportingParty *TradeParty2 `xml:"RptgPty"`

	// Agent which registers the currency control contract.
	RegistrationAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RegnAgt"`

	// Identifies the requested registered contract.
	RegisteredContractIdentification *Max35Text `xml:"RegdCtrctId"`

	// Defines the criteria to be returned in the statement in response to the request.
	ReturnCriteria *ContractRegistrationStatementCriteria1 `xml:"RtrCrit,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *ContractRegistrationStatementRequest1) SetStatementRequestIdentification(value string) {
	c.StatementRequestIdentification = (*Max35Text)(&value)
}

func (c *ContractRegistrationStatementRequest1) AddReportingPeriod() *ReportingPeriod1 {
	c.ReportingPeriod = new(ReportingPeriod1)
	return c.ReportingPeriod
}

func (c *ContractRegistrationStatementRequest1) AddReportingParty() *TradeParty2 {
	c.ReportingParty = new(TradeParty2)
	return c.ReportingParty
}

func (c *ContractRegistrationStatementRequest1) AddRegistrationAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.RegistrationAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.RegistrationAgent
}

func (c *ContractRegistrationStatementRequest1) SetRegisteredContractIdentification(value string) {
	c.RegisteredContractIdentification = (*Max35Text)(&value)
}

func (c *ContractRegistrationStatementRequest1) AddReturnCriteria() *ContractRegistrationStatementCriteria1 {
	c.ReturnCriteria = new(ContractRegistrationStatementCriteria1)
	return c.ReturnCriteria
}

func (c *ContractRegistrationStatementRequest1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
