package model

// Sequence of terminal management actions to be performed by a point of interaction (POI).
type ManagementPlan3 struct {

	// Identification of the point of interaction (POI) for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the management plan.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId"`

	// Data set related to the sequence of actions to be performed by a point of interaction (POI).
	DataSet []*TerminalManagementDataSet10 `xml:"DataSet"`
}

func (m *ManagementPlan3) AddPOIIdentification() *GenericIdentification35 {
	m.POIIdentification = new(GenericIdentification35)
	return m.POIIdentification
}

func (m *ManagementPlan3) AddTerminalManagerIdentification() *GenericIdentification35 {
	m.TerminalManagerIdentification = new(GenericIdentification35)
	return m.TerminalManagerIdentification
}

func (m *ManagementPlan3) AddDataSet() *TerminalManagementDataSet10 {
	newValue := new(TerminalManagementDataSet10)
	m.DataSet = append(m.DataSet, newValue)
	return newValue
}
