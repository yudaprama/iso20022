package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope28 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
	InvestmentPlan *InvestmentPlan13 `xml:"InvstmtPlan"`
}

func (m *ModificationScope28) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope28) AddInvestmentPlan() *InvestmentPlan13 {
	m.InvestmentPlan = new(InvestmentPlan13)
	return m.InvestmentPlan
}
