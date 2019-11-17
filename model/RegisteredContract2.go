package model

// Document that a user must file with an authorized servicer for each contract that involves foreign currency transactions with non residents.
type RegisteredContract2 struct {

	// Unique and unambiguous identification of the registered contract closure.
	RegisteredContractClosureIdentification *Max35Text `xml:"RegdCtrctClsrId"`

	// Party who registered the currency control contract.
	ReportingParty *TradeParty2 `xml:"RptgPty"`

	// Agent who registered the currency control contract.
	RegistrationAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RegnAgt"`

	// Original registered contract identification.
	OriginalRegisteredContract *DocumentIdentification29 `xml:"OrgnlRegdCtrct"`

	// Priority of the registered contract closure.
	Priority *Priority2Code `xml:"Prty"`

	// Reason of the closure.
	ClosureReason *ContractClosureReason1Choice `xml:"ClsrRsn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RegisteredContract2) SetRegisteredContractClosureIdentification(value string) {
	r.RegisteredContractClosureIdentification = (*Max35Text)(&value)
}

func (r *RegisteredContract2) AddReportingParty() *TradeParty2 {
	r.ReportingParty = new(TradeParty2)
	return r.ReportingParty
}

func (r *RegisteredContract2) AddRegistrationAgent() *BranchAndFinancialInstitutionIdentification5 {
	r.RegistrationAgent = new(BranchAndFinancialInstitutionIdentification5)
	return r.RegistrationAgent
}

func (r *RegisteredContract2) AddOriginalRegisteredContract() *DocumentIdentification29 {
	r.OriginalRegisteredContract = new(DocumentIdentification29)
	return r.OriginalRegisteredContract
}

func (r *RegisteredContract2) SetPriority(value string) {
	r.Priority = (*Priority2Code)(&value)
}

func (r *RegisteredContract2) AddClosureReason() *ContractClosureReason1Choice {
	r.ClosureReason = new(ContractClosureReason1Choice)
	return r.ClosureReason
}

func (r *RegisteredContract2) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}
