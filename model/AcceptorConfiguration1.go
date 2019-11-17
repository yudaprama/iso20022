package model

// Acceptor configuration to be downloaded from the terminal management system.
type AcceptorConfiguration1 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the acceptor parameters.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId"`

	// Data set containing the acceptor parameters of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet3 `xml:"DataSet"`
}

func (a *AcceptorConfiguration1) AddPOIIdentification() *GenericIdentification35 {
	a.POIIdentification = new(GenericIdentification35)
	return a.POIIdentification
}

func (a *AcceptorConfiguration1) AddTerminalManagerIdentification() *GenericIdentification35 {
	a.TerminalManagerIdentification = new(GenericIdentification35)
	return a.TerminalManagerIdentification
}

func (a *AcceptorConfiguration1) AddDataSet() *TerminalManagementDataSet3 {
	newValue := new(TerminalManagementDataSet3)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}
