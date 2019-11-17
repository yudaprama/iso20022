package model

// Sequence of terminal management actions to be performed by a point of interaction (POI).
type ManagementPlan1 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the management plan.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId"`

	// Data set related to the sequence of actions to be performed by a point of interaction (POI).
	DataSet []*TerminalManagementDataSet2 `xml:"DataSet"`
}

func (m *ManagementPlan1) AddPOIIdentification() *GenericIdentification35 {
	m.POIIdentification = new(GenericIdentification35)
	return m.POIIdentification
}

func (m *ManagementPlan1) AddTerminalManagerIdentification() *GenericIdentification35 {
	m.TerminalManagerIdentification = new(GenericIdentification35)
	return m.TerminalManagerIdentification
}

func (m *ManagementPlan1) AddDataSet() *TerminalManagementDataSet2 {
	newValue := new(TerminalManagementDataSet2)
	m.DataSet = append(m.DataSet, newValue)
	return newValue
}
