package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope16 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
	InvestmentPlan *InvestmentPlan7 `xml:"InvstmtPlan"`
}

func (m *ModificationScope16) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope16) AddInvestmentPlan() *InvestmentPlan7 {
	m.InvestmentPlan = new(InvestmentPlan7)
	return m.InvestmentPlan
}
