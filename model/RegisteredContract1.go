package model

// Currency control registered contract amendment details.
type RegisteredContract1 struct {

	// Unique and unambiguous identification of the  contract registration amendment.
	ContractRegistrationAmendmentIdentification *Max35Text `xml:"CtrctRegnAmdmntId"`

	// Party registering the currency control contract.
	ReportingParty *TradeParty2 `xml:"RptgPty"`

	// Agent which registers the currency control contract.
	RegistrationAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RegnAgt"`

	// Amendment details applied on one or several registered contracts.
	RegisteredContractAmendment []*RegisteredContract3 `xml:"RegdCtrctAmdmnt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RegisteredContract1) SetContractRegistrationAmendmentIdentification(value string) {
	r.ContractRegistrationAmendmentIdentification = (*Max35Text)(&value)
}

func (r *RegisteredContract1) AddReportingParty() *TradeParty2 {
	r.ReportingParty = new(TradeParty2)
	return r.ReportingParty
}

func (r *RegisteredContract1) AddRegistrationAgent() *BranchAndFinancialInstitutionIdentification5 {
	r.RegistrationAgent = new(BranchAndFinancialInstitutionIdentification5)
	return r.RegistrationAgent
}

func (r *RegisteredContract1) AddRegisteredContractAmendment() *RegisteredContract3 {
	newValue := new(RegisteredContract3)
	r.RegisteredContractAmendment = append(r.RegisteredContractAmendment, newValue)
	return newValue
}

func (r *RegisteredContract1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}
