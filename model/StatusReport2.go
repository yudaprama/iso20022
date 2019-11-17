package model

// Status of the acceptor system containing the identification of the POI, its components and their installed versions.
type StatusReport2 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId"`

	// Identification of the terminal management system (TMS) to contact for the maintenance.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId,omitempty"`

	// Data related to a status report of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet4 `xml:"DataSet"`
}

func (s *StatusReport2) AddPOIIdentification() *GenericIdentification35 {
	s.POIIdentification = new(GenericIdentification35)
	return s.POIIdentification
}

func (s *StatusReport2) AddTerminalManagerIdentification() *GenericIdentification35 {
	s.TerminalManagerIdentification = new(GenericIdentification35)
	return s.TerminalManagerIdentification
}

func (s *StatusReport2) AddDataSet() *TerminalManagementDataSet4 {
	newValue := new(TerminalManagementDataSet4)
	s.DataSet = append(s.DataSet, newValue)
	return newValue
}
