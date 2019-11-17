package model

// Acceptor configuration to be downloaded from the terminal management system.
type AcceptorConfiguration5 struct {

	// Identification of the terminal management system (TMS) sending the acceptor parameters.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId"`

	// Data set containing the acceptor parameters of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet19 `xml:"DataSet"`
}

func (a *AcceptorConfiguration5) AddTerminalManagerIdentification() *GenericIdentification71 {
	a.TerminalManagerIdentification = new(GenericIdentification71)
	return a.TerminalManagerIdentification
}

func (a *AcceptorConfiguration5) AddDataSet() *TerminalManagementDataSet19 {
	newValue := new(TerminalManagementDataSet19)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}
