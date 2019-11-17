package model

// Provides securities proceeds information.
type SecuritiesProceeds1 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification7 `xml:"SctyId"`

	// The quantity of financial instruments that is posted.
	PostingQuantity *UnitOrFaceAmount1Choice `xml:"PstngQty"`

	// Provides information about the account that is debited/credited.
	AccountDetails []*SecuritiesAccount10 `xml:"AcctDtls"`

	// Provides reconciliation information.
	ReconciliationDetails *Max350Text `xml:"RcncltnDtls,omitempty"`
}

func (s *SecuritiesProceeds1) AddSecurityIdentification() *SecurityIdentification7 {
	s.SecurityIdentification = new(SecurityIdentification7)
	return s.SecurityIdentification
}

func (s *SecuritiesProceeds1) AddPostingQuantity() *UnitOrFaceAmount1Choice {
	s.PostingQuantity = new(UnitOrFaceAmount1Choice)
	return s.PostingQuantity
}

func (s *SecuritiesProceeds1) AddAccountDetails() *SecuritiesAccount10 {
	newValue := new(SecuritiesAccount10)
	s.AccountDetails = append(s.AccountDetails, newValue)
	return newValue
}

func (s *SecuritiesProceeds1) SetReconciliationDetails(value string) {
	s.ReconciliationDetails = (*Max350Text)(&value)
}
