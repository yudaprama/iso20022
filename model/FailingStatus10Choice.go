package model

// Choice of failing status.
type FailingStatus10Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the failing status.
	Reason []*FailingReason8 `xml:"Rsn"`
}

func (f *FailingStatus10Choice) SetNoSpecifiedReason(value string) {
	f.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (f *FailingStatus10Choice) AddReason() *FailingReason8 {
	newValue := new(FailingReason8)
	f.Reason = append(f.Reason, newValue)
	return newValue
}
