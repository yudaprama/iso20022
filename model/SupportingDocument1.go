package model

// Currency control document supporting the contract registration.
type SupportingDocument1 struct {

	// Unique and unambiguous identification of the supporting document.
	SupportingDocumentIdentification *Max35Text `xml:"SpprtgDocId"`

	// Unique identification of the original query message.
	OriginalRequestIdentification *Max35Text `xml:"OrgnlReqId,omitempty"`

	// Unique identification of the certificate for which the supporting document is provided.
	Certificate *DocumentIdentification28 `xml:"Cert"`

	// Party that legally owns the cash account.
	AccountOwner *PartyIdentification77 `xml:"AcctOwnr"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr"`

	// Amendment indicator details.
	Amendment *DocumentAmendment1 `xml:"Amdmnt,omitempty"`

	// Reference of the registered contract or the underlying contract for the supporting documents.
	ContractReference *ContractRegistrationReference1Choice `xml:"CtrctRef"`

	// Individual entry of the supporting document.
	Entry []*SupportingDocumentEntry1 `xml:"Ntry"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SupportingDocument1) SetSupportingDocumentIdentification(value string) {
	s.SupportingDocumentIdentification = (*Max35Text)(&value)
}

func (s *SupportingDocument1) SetOriginalRequestIdentification(value string) {
	s.OriginalRequestIdentification = (*Max35Text)(&value)
}

func (s *SupportingDocument1) AddCertificate() *DocumentIdentification28 {
	s.Certificate = new(DocumentIdentification28)
	return s.Certificate
}

func (s *SupportingDocument1) AddAccountOwner() *PartyIdentification77 {
	s.AccountOwner = new(PartyIdentification77)
	return s.AccountOwner
}

func (s *SupportingDocument1) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	s.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return s.AccountServicer
}

func (s *SupportingDocument1) AddAmendment() *DocumentAmendment1 {
	s.Amendment = new(DocumentAmendment1)
	return s.Amendment
}

func (s *SupportingDocument1) AddContractReference() *ContractRegistrationReference1Choice {
	s.ContractReference = new(ContractRegistrationReference1Choice)
	return s.ContractReference
}

func (s *SupportingDocument1) AddEntry() *SupportingDocumentEntry1 {
	newValue := new(SupportingDocumentEntry1)
	s.Entry = append(s.Entry, newValue)
	return newValue
}

func (s *SupportingDocument1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
