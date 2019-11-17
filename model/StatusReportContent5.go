package model

// Content of the status report.
type StatusReportContent5 struct {

	// Capabilities of the POI (Point Of Interaction) performing the status report.
	POICapabilities *PointOfInteractionCapabilities6 `xml:"POICpblties,omitempty"`

	// Data related to a component of the POI (Point Of Interaction) performing the status report.
	POIComponent []*PointOfInteractionComponent6 `xml:"POICmpnt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during transactions.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// System date time of the point of interaction (POI) sending the status report.
	POIDateTime *ISODateTime `xml:"POIDtTm"`

	// Request the terminal management system to answer with the identified data set.
	DataSetRequired *TerminalManagementDataSet17 `xml:"DataSetReqrd,omitempty"`

	// Result of an individual terminal management action by the point of interaction.
	Event []*TMSEvent4 `xml:"Evt,omitempty"`

	// Error log of the point of interaction since the last status report.
	Errors []*Max140Text `xml:"Errs,omitempty"`
}

func (s *StatusReportContent5) AddPOICapabilities() *PointOfInteractionCapabilities6 {
	s.POICapabilities = new(PointOfInteractionCapabilities6)
	return s.POICapabilities
}

func (s *StatusReportContent5) AddPOIComponent() *PointOfInteractionComponent6 {
	newValue := new(PointOfInteractionComponent6)
	s.POIComponent = append(s.POIComponent, newValue)
	return newValue
}

func (s *StatusReportContent5) SetAttendanceContext(value string) {
	s.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (s *StatusReportContent5) SetPOIDateTime(value string) {
	s.POIDateTime = (*ISODateTime)(&value)
}

func (s *StatusReportContent5) AddDataSetRequired() *TerminalManagementDataSet17 {
	s.DataSetRequired = new(TerminalManagementDataSet17)
	return s.DataSetRequired
}

func (s *StatusReportContent5) AddEvent() *TMSEvent4 {
	newValue := new(TMSEvent4)
	s.Event = append(s.Event, newValue)
	return newValue
}

func (s *StatusReportContent5) AddErrors(value string) {
	s.Errors = append(s.Errors, (*Max140Text)(&value))
}
