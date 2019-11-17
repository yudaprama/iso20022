package model

// Content of the status report.
type StatusReportContent4 struct {

	// Capabilities of the POI (Point Of Interaction) performing the status report.
	POICapabilities *PointOfInteractionCapabilities3 `xml:"POICpblties,omitempty"`

	// Data related to a component of the POI (Point Of Interaction) performing the status report.
	POIComponent []*PointOfInteractionComponent5 `xml:"POICmpnt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during transactions.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// System date time of the point of interaction (POI) sending the status report.
	POIDateTime *ISODateTime `xml:"POIDtTm"`

	// Request the terminal management system to answer with the identified data set.
	DataSetRequired *TerminalManagementDataSet12 `xml:"DataSetReqrd,omitempty"`

	// Result of an individual terminal management action by the point of interaction.
	Event []*TMSEvent3 `xml:"Evt,omitempty"`

	// Error log of the point of interaction since the last status report.
	Errors []*Max140Text `xml:"Errs,omitempty"`
}

func (s *StatusReportContent4) AddPOICapabilities() *PointOfInteractionCapabilities3 {
	s.POICapabilities = new(PointOfInteractionCapabilities3)
	return s.POICapabilities
}

func (s *StatusReportContent4) AddPOIComponent() *PointOfInteractionComponent5 {
	newValue := new(PointOfInteractionComponent5)
	s.POIComponent = append(s.POIComponent, newValue)
	return newValue
}

func (s *StatusReportContent4) SetAttendanceContext(value string) {
	s.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (s *StatusReportContent4) SetPOIDateTime(value string) {
	s.POIDateTime = (*ISODateTime)(&value)
}

func (s *StatusReportContent4) AddDataSetRequired() *TerminalManagementDataSet12 {
	s.DataSetRequired = new(TerminalManagementDataSet12)
	return s.DataSetRequired
}

func (s *StatusReportContent4) AddEvent() *TMSEvent3 {
	newValue := new(TMSEvent3)
	s.Event = append(s.Event, newValue)
	return newValue
}

func (s *StatusReportContent4) AddErrors(value string) {
	s.Errors = append(s.Errors, (*Max140Text)(&value))
}
