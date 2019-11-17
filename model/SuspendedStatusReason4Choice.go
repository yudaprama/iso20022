package model

// Reason for a suspended status.
type SuspendedStatusReason4Choice struct {

	// No reason available or to report for the suspended status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the suspended status, expressed as a code.
	ReasonDetails []*SuspendedStatusReason4 `xml:"RsnDtls"`
}

func (s *SuspendedStatusReason4Choice) SetNoSpecifiedReason(value string) {
	s.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (s *SuspendedStatusReason4Choice) AddReasonDetails() *SuspendedStatusReason4 {
	newValue := new(SuspendedStatusReason4)
	s.ReasonDetails = append(s.ReasonDetails, newValue)
	return newValue
}
