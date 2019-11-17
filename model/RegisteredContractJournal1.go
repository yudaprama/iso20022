package model

// Document that a user must file with an authorised servicer for each contract that involves foreign currency transactions with non residents.
type RegisteredContractJournal1 struct {

	// Agent which registers the currency control contract.
	RegistrationAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RegnAgt"`

	// Unique registered contract identification.
	UniqueIdentification *DocumentIdentification28 `xml:"UnqId,omitempty"`

	// Date of closure of the registered contract.
	ClosureDate *ISODate `xml:"ClsrDt"`

	// Reason of closure
	//
	// TBD - codes to be defined.
	ClosureReason *ContractClosureReason1Choice `xml:"ClsrRsn"`
}

func (r *RegisteredContractJournal1) AddRegistrationAgent() *BranchAndFinancialInstitutionIdentification5 {
	r.RegistrationAgent = new(BranchAndFinancialInstitutionIdentification5)
	return r.RegistrationAgent
}

func (r *RegisteredContractJournal1) AddUniqueIdentification() *DocumentIdentification28 {
	r.UniqueIdentification = new(DocumentIdentification28)
	return r.UniqueIdentification
}

func (r *RegisteredContractJournal1) SetClosureDate(value string) {
	r.ClosureDate = (*ISODate)(&value)
}

func (r *RegisteredContractJournal1) AddClosureReason() *ContractClosureReason1Choice {
	r.ClosureReason = new(ContractClosureReason1Choice)
	return r.ClosureReason
}
