package model

// Content of the status report.
type StatusReportContent2 struct {

	// Capabilities of the POI performing the status report.
	POICapabilities *PointOfInteractionCapabilities1 `xml:"POICpblties,omitempty"`

	// Data related to a component of the POI performing the status report.
	POIComponent []*PointOfInteractionComponent3 `xml:"POICmpnt,omitempty"`

	// Human attendance at the POI location during transactions.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// System date time of the point of interaction (POI) sending the status report.
	POIDateTime *ISODateTime `xml:"POIDtTm"`

	// Request the terminal management system to answer with the identified data set.
	DataSetRequired *TerminalManagementDataSet7 `xml:"DataSetReqrd,omitempty"`

	// Result of an individual terminal management action by the point of interaction.
	Event []*TMSEvent2 `xml:"Evt,omitempty"`

	// Error log of the point of interaction since the last status report.
	Errors *Max140Text `xml:"Errs,omitempty"`
}

func (s *StatusReportContent2) AddPOICapabilities() *PointOfInteractionCapabilities1 {
	s.POICapabilities = new(PointOfInteractionCapabilities1)
	return s.POICapabilities
}

func (s *StatusReportContent2) AddPOIComponent() *PointOfInteractionComponent3 {
	newValue := new(PointOfInteractionComponent3)
	s.POIComponent = append(s.POIComponent, newValue)
	return newValue
}

func (s *StatusReportContent2) SetAttendanceContext(value string) {
	s.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (s *StatusReportContent2) SetPOIDateTime(value string) {
	s.POIDateTime = (*ISODateTime)(&value)
}

func (s *StatusReportContent2) AddDataSetRequired() *TerminalManagementDataSet7 {
	s.DataSetRequired = new(TerminalManagementDataSet7)
	return s.DataSetRequired
}

func (s *StatusReportContent2) AddEvent() *TMSEvent2 {
	newValue := new(TMSEvent2)
	s.Event = append(s.Event, newValue)
	return newValue
}

func (s *StatusReportContent2) SetErrors(value string) {
	s.Errors = (*Max140Text)(&value)
}
