package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope8 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
	InvestmentPlan *InvestmentPlan5 `xml:"InvstmtPlan"`
}

func (m *ModificationScope8) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope8) AddInvestmentPlan() *InvestmentPlan5 {
	m.InvestmentPlan = new(InvestmentPlan5)
	return m.InvestmentPlan
}
