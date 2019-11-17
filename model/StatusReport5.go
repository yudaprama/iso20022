package model

// Status of the acceptor system containing the identification of the POI (Point Of Interaction), its components and their installed versions.
type StatusReport5 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification71 `xml:"POIId"`

	// Identification of the terminal management system (TMS) to contact for the maintenance.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId"`

	// Data related to a status report of a point of interaction (POI).
	DataSet *TerminalManagementDataSet16 `xml:"DataSet"`
}

func (s *StatusReport5) AddPOIIdentification() *GenericIdentification71 {
	s.POIIdentification = new(GenericIdentification71)
	return s.POIIdentification
}

func (s *StatusReport5) AddTerminalManagerIdentification() *GenericIdentification71 {
	s.TerminalManagerIdentification = new(GenericIdentification71)
	return s.TerminalManagerIdentification
}

func (s *StatusReport5) AddDataSet() *TerminalManagementDataSet16 {
	s.DataSet = new(TerminalManagementDataSet16)
	return s.DataSet
}
