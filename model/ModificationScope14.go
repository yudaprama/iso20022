package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope14 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Detailed information about the party profile information.
	InvestorProfileValidation *PartyProfileInformation2 `xml:"InvstrPrflVldtn"`
}

func (m *ModificationScope14) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope14) AddInvestorProfileValidation() *PartyProfileInformation2 {
	m.InvestorProfileValidation = new(PartyProfileInformation2)
	return m.InvestorProfileValidation
}
