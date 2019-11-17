package model

// Acceptor configuration to be downloaded from the terminal management system.
type AcceptorConfiguration4 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification71 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the acceptor parameters.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId"`

	// Data set containing the acceptor parameters of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet14 `xml:"DataSet"`
}

func (a *AcceptorConfiguration4) AddPOIIdentification() *GenericIdentification71 {
	a.POIIdentification = new(GenericIdentification71)
	return a.POIIdentification
}

func (a *AcceptorConfiguration4) AddTerminalManagerIdentification() *GenericIdentification71 {
	a.TerminalManagerIdentification = new(GenericIdentification71)
	return a.TerminalManagerIdentification
}

func (a *AcceptorConfiguration4) AddDataSet() *TerminalManagementDataSet14 {
	newValue := new(TerminalManagementDataSet14)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}
