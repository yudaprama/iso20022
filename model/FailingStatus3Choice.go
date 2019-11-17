package model

// Choice of failing status.
type FailingStatus3Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the FailingStatus.
	Reason []*FailingReason3 `xml:"Rsn"`
}

func (f *FailingStatus3Choice) SetNoSpecifiedReason(value string) {
	f.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (f *FailingStatus3Choice) AddReason() *FailingReason3 {
	newValue := new(FailingReason3)
	f.Reason = append(f.Reason, newValue)
	return newValue
}
