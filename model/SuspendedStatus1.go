package model

// Status is suspended.
type SuspendedStatus1 struct {

	// Indicates that there is no reason available or to report.
	NoReason *NoReasonCode `xml:"NoRsn"`

	// Reason for a suspended status in the report.
	Reason *SuspendedStatusReason1 `xml:"Rsn"`

	// Proprietary identification of a reason for a suspended status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (s *SuspendedStatus1) SetNoReason(value string) {
	s.NoReason = (*NoReasonCode)(&value)
}

func (s *SuspendedStatus1) AddReason() *SuspendedStatusReason1 {
	s.Reason = new(SuspendedStatusReason1)
	return s.Reason
}

func (s *SuspendedStatus1) AddDataSourceScheme() *GenericIdentification1 {
	s.DataSourceScheme = new(GenericIdentification1)
	return s.DataSourceScheme
}
