package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope21 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Information about the investment account ownership with respect to new issue allocation for a hedge fund.
	IssueAllocation *NewIssueAllocation2 `xml:"IsseAllcn"`
}

func (m *ModificationScope21) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope21) AddIssueAllocation() *NewIssueAllocation2 {
	m.IssueAllocation = new(NewIssueAllocation2)
	return m.IssueAllocation
}
