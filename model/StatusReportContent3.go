package model

// Content of the status report.
type StatusReportContent3 struct {

	// Capabilities of the POI (Point Of Interaction) performing the status report.
	POICapabilities *PointOfInteractionCapabilities2 `xml:"POICpblties,omitempty"`

	// Data related to a component of the POI (Point Of Interaction) performing the status report.
	POIComponent []*PointOfInteractionComponent4 `xml:"POICmpnt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during transactions.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// System date time of the point of interaction (POI) sending the status report.
	POIDateTime *ISODateTime `xml:"POIDtTm"`

	// Request the terminal management system to answer with the identified data set.
	DataSetRequired *TerminalManagementDataSet8 `xml:"DataSetReqrd,omitempty"`

	// Result of an individual terminal management action by the point of interaction.
	Event []*TMSEvent2 `xml:"Evt,omitempty"`

	// Error log of the point of interaction since the last status report.
	Errors []*Max140Text `xml:"Errs,omitempty"`
}

func (s *StatusReportContent3) AddPOICapabilities() *PointOfInteractionCapabilities2 {
	s.POICapabilities = new(PointOfInteractionCapabilities2)
	return s.POICapabilities
}

func (s *StatusReportContent3) AddPOIComponent() *PointOfInteractionComponent4 {
	newValue := new(PointOfInteractionComponent4)
	s.POIComponent = append(s.POIComponent, newValue)
	return newValue
}

func (s *StatusReportContent3) SetAttendanceContext(value string) {
	s.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (s *StatusReportContent3) SetPOIDateTime(value string) {
	s.POIDateTime = (*ISODateTime)(&value)
}

func (s *StatusReportContent3) AddDataSetRequired() *TerminalManagementDataSet8 {
	s.DataSetRequired = new(TerminalManagementDataSet8)
	return s.DataSetRequired
}

func (s *StatusReportContent3) AddEvent() *TMSEvent2 {
	newValue := new(TMSEvent2)
	s.Event = append(s.Event, newValue)
	return newValue
}

func (s *StatusReportContent3) AddErrors(value string) {
	s.Errors = append(s.Errors, (*Max140Text)(&value))
}
