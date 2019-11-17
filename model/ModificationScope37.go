package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope37 struct {

	// Type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
	InvestmentPlan *InvestmentPlan15 `xml:"InvstmtPlan"`
}

func (m *ModificationScope37) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope37) AddInvestmentPlan() *InvestmentPlan15 {
	m.InvestmentPlan = new(InvestmentPlan15)
	return m.InvestmentPlan
}
