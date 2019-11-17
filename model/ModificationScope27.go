package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope27 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Detailed information about the party profile information.
	InvestorProfileValidation *PartyProfileInformation5 `xml:"InvstrPrflVldtn"`
}

func (m *ModificationScope27) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope27) AddInvestorProfileValidation() *PartyProfileInformation5 {
	m.InvestorProfileValidation = new(PartyProfileInformation5)
	return m.InvestorProfileValidation
}
