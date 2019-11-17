package model

// Acceptor configuration to be downloaded from the terminal management system.
type AcceptorConfiguration2 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the acceptor parameters.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId"`

	// Data set containing the acceptor parameters of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet6 `xml:"DataSet"`
}

func (a *AcceptorConfiguration2) AddPOIIdentification() *GenericIdentification35 {
	a.POIIdentification = new(GenericIdentification35)
	return a.POIIdentification
}

func (a *AcceptorConfiguration2) AddTerminalManagerIdentification() *GenericIdentification35 {
	a.TerminalManagerIdentification = new(GenericIdentification35)
	return a.TerminalManagerIdentification
}

func (a *AcceptorConfiguration2) AddDataSet() *TerminalManagementDataSet6 {
	newValue := new(TerminalManagementDataSet6)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}
