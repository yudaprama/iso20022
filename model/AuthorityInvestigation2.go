package model

// Specifies the requested authority investigation information details.
type AuthorityInvestigation2 struct {

	// Identifies the type requested information as a code.
	Type *AuthorityRequestType1 `xml:"Tp"`

	// Identifies the roles the customer plays in the requested information.
	InvestigatedRoles *InvestigatedParties1Choice `xml:"InvstgtdRoles"`

	// Specifies the additional investigated parties.
	AdditionalInvestigatedParties *InvestigatedParties1Choice `xml:"AddtlInvstgtdPties,omitempty"`

	// Additional information, in free text form, to complement the requested information.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`
}

func (a *AuthorityInvestigation2) AddType() *AuthorityRequestType1 {
	a.Type = new(AuthorityRequestType1)
	return a.Type
}

func (a *AuthorityInvestigation2) AddInvestigatedRoles() *InvestigatedParties1Choice {
	a.InvestigatedRoles = new(InvestigatedParties1Choice)
	return a.InvestigatedRoles
}

func (a *AuthorityInvestigation2) AddAdditionalInvestigatedParties() *InvestigatedParties1Choice {
	a.AdditionalInvestigatedParties = new(InvestigatedParties1Choice)
	return a.AdditionalInvestigatedParties
}

func (a *AuthorityInvestigation2) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max500Text)(&value)
}
