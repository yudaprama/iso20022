package model

// Currency control contract registration details.
type ContractRegistration1 struct {

	// Unique and unambiguous identification of the contract registration.
	ContractRegistrationIdentification *Max35Text `xml:"CtrctRegnId"`

	// Party registering the currency control contract.
	ReportingParty *TradeParty2 `xml:"RptgPty"`

	// Agent which registers the currency control contract.
	RegistrationAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RegnAgt"`

	// Details about the opening of the contract registration.
	ContractRegistrationOpening []*ContractRegistration2 `xml:"CtrctRegnOpng"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *ContractRegistration1) SetContractRegistrationIdentification(value string) {
	c.ContractRegistrationIdentification = (*Max35Text)(&value)
}

func (c *ContractRegistration1) AddReportingParty() *TradeParty2 {
	c.ReportingParty = new(TradeParty2)
	return c.ReportingParty
}

func (c *ContractRegistration1) AddRegistrationAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.RegistrationAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.RegistrationAgent
}

func (c *ContractRegistration1) AddContractRegistrationOpening() *ContractRegistration2 {
	newValue := new(ContractRegistration2)
	c.ContractRegistrationOpening = append(c.ContractRegistrationOpening, newValue)
	return newValue
}

func (c *ContractRegistration1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
