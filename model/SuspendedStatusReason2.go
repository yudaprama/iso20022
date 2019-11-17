package model

// Identification of the reason for the suspended status.
type SuspendedStatusReason2 struct {

	// Reason for the suspended status.
	Reason *SuspendedStatusReason3Code `xml:"Rsn"`

	// Reason for the suspended status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the suspended status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Additional information about the suspended status reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (s *SuspendedStatusReason2) SetReason(value string) {
	s.Reason = (*SuspendedStatusReason3Code)(&value)
}

func (s *SuspendedStatusReason2) SetExtendedReason(value string) {
	s.ExtendedReason = (*Extended350Code)(&value)
}

func (s *SuspendedStatusReason2) AddDataSourceScheme() *GenericIdentification1 {
	s.DataSourceScheme = new(GenericIdentification1)
	return s.DataSourceScheme
}

func (s *SuspendedStatusReason2) SetAdditionalInformation(value string) {
	s.AdditionalInformation = (*Max350Text)(&value)
}
