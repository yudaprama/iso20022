package model

// Status is suspended.
type SuspendedStatus2 struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the suspended status.
	ReasonDetails []*SuspendedStatusReason2 `xml:"RsnDtls"`
}

func (s *SuspendedStatus2) SetNoSpecifiedReason(value string) {
	s.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (s *SuspendedStatus2) AddReasonDetails() *SuspendedStatusReason2 {
	newValue := new(SuspendedStatusReason2)
	s.ReasonDetails = append(s.ReasonDetails, newValue)
	return newValue
}
