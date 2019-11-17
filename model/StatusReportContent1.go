package model

// Content of the status report.
type StatusReportContent1 struct {

	// Capabilities of the POI performing the status report.
	POICapabilities *PointOfInteractionCapabilities1 `xml:"POICpblties,omitempty"`

	// Data related to a component of the POI performing the status report.
	POIComponent []*PointOfInteractionComponent2 `xml:"POICmpnt,omitempty"`

	// Human attendance at the POI location during transactions.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// System date time of the point of interaction (POI) sending the status report.
	POIDateTime *ISODateTime `xml:"POIDtTm"`

	// Request the terminal management system to answer with the identified data set.
	DataSetRequired *DataSetIdentification2 `xml:"DataSetReqrd,omitempty"`

	// Result of an individual terminal management action by the point of interaction.
	Event []*TMSEvent1 `xml:"Evt,omitempty"`

	// Error log of the point of interaction since the last status report.
	Errors *Max140Text `xml:"Errs,omitempty"`
}

func (s *StatusReportContent1) AddPOICapabilities() *PointOfInteractionCapabilities1 {
	s.POICapabilities = new(PointOfInteractionCapabilities1)
	return s.POICapabilities
}

func (s *StatusReportContent1) AddPOIComponent() *PointOfInteractionComponent2 {
	newValue := new(PointOfInteractionComponent2)
	s.POIComponent = append(s.POIComponent, newValue)
	return newValue
}

func (s *StatusReportContent1) SetAttendanceContext(value string) {
	s.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (s *StatusReportContent1) SetPOIDateTime(value string) {
	s.POIDateTime = (*ISODateTime)(&value)
}

func (s *StatusReportContent1) AddDataSetRequired() *DataSetIdentification2 {
	s.DataSetRequired = new(DataSetIdentification2)
	return s.DataSetRequired
}

func (s *StatusReportContent1) AddEvent() *TMSEvent1 {
	newValue := new(TMSEvent1)
	s.Event = append(s.Event, newValue)
	return newValue
}

func (s *StatusReportContent1) SetErrors(value string) {
	s.Errors = (*Max140Text)(&value)
}
