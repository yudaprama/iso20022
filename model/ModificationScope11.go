package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope11 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Detailed information about the party profile information.
	InvestorProfileValidation *PartyProfileInformation1 `xml:"InvstrPrflVldtn"`
}

func (m *ModificationScope11) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope11) AddInvestorProfileValidation() *PartyProfileInformation1 {
	m.InvestorProfileValidation = new(PartyProfileInformation1)
	return m.InvestorProfileValidation
}
