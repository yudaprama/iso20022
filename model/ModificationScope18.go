package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope18 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
	InvestmentPlan *InvestmentPlan8 `xml:"InvstmtPlan"`
}

func (m *ModificationScope18) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope18) AddInvestmentPlan() *InvestmentPlan8 {
	m.InvestmentPlan = new(InvestmentPlan8)
	return m.InvestmentPlan
}
