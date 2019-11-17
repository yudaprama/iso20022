package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope32 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Information used to determine fees and types of operations that can be carried out on the account.
	InvestorProfile *InvestorProfile1 `xml:"InvstrPrfl"`
}

func (m *ModificationScope32) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope32) AddInvestorProfile() *InvestorProfile1 {
	m.InvestorProfile = new(InvestorProfile1)
	return m.InvestorProfile
}
