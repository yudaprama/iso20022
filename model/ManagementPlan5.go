package model

// Sequence of terminal management actions to be performed by a point of interaction (POI).
type ManagementPlan5 struct {

	// Identification of the point of interaction (POI) for terminal management.
	POIIdentification *GenericIdentification71 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the management plan.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId"`

	// Data set related to the sequence of actions to be performed by a point of interaction (POI).
	DataSet *TerminalManagementDataSet18 `xml:"DataSet"`
}

func (m *ManagementPlan5) AddPOIIdentification() *GenericIdentification71 {
	m.POIIdentification = new(GenericIdentification71)
	return m.POIIdentification
}

func (m *ManagementPlan5) AddTerminalManagerIdentification() *GenericIdentification71 {
	m.TerminalManagerIdentification = new(GenericIdentification71)
	return m.TerminalManagerIdentification
}

func (m *ManagementPlan5) AddDataSet() *TerminalManagementDataSet18 {
	m.DataSet = new(TerminalManagementDataSet18)
	return m.DataSet
}
