package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope19 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Detailed information about the party profile information.
	InvestorProfileValidation *PartyProfileInformation3 `xml:"InvstrPrflVldtn"`
}

func (m *ModificationScope19) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope19) AddInvestorProfileValidation() *PartyProfileInformation3 {
	m.InvestorProfileValidation = new(PartyProfileInformation3)
	return m.InvestorProfileValidation
}
