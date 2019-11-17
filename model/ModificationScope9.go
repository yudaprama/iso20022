package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope9 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Information about the investment account ownership with respect to new issue allocation for a hedge fund.
	IssueAllocation *NewIssueAllocation1 `xml:"IsseAllcn"`
}

func (m *ModificationScope9) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope9) AddIssueAllocation() *NewIssueAllocation1 {
	m.IssueAllocation = new(NewIssueAllocation1)
	return m.IssueAllocation
}
