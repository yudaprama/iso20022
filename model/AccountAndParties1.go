package model

// Identifies the account as the search criteria for the financial institution to do the investigation.
type AccountAndParties1 struct {

	// Specifies the account for the investigation.
	Identification *CashAccount25 `xml:"Id"`

	// Specifies the investigated parties related to the account such as the owner, beneficiary, signatory or any party playing a role in that account for which the investigation needs to be done.
	InvestigatedParties *InvestigatedParties1Choice `xml:"InvstgtdPties"`

	// Identifies the authority request type as a code.
	AuthorityRequestType []*AuthorityRequestType1 `xml:"AuthrtyReqTp"`
}

func (a *AccountAndParties1) AddIdentification() *CashAccount25 {
	a.Identification = new(CashAccount25)
	return a.Identification
}

func (a *AccountAndParties1) AddInvestigatedParties() *InvestigatedParties1Choice {
	a.InvestigatedParties = new(InvestigatedParties1Choice)
	return a.InvestigatedParties
}

func (a *AccountAndParties1) AddAuthorityRequestType() *AuthorityRequestType1 {
	newValue := new(AuthorityRequestType1)
	a.AuthorityRequestType = append(a.AuthorityRequestType, newValue)
	return newValue
}
