package model

// Choice of failing status.
type FailingStatus9Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the failing status.
	Reason []*FailingReason7 `xml:"Rsn"`
}

func (f *FailingStatus9Choice) SetNoSpecifiedReason(value string) {
	f.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (f *FailingStatus9Choice) AddReason() *FailingReason7 {
	newValue := new(FailingReason7)
	f.Reason = append(f.Reason, newValue)
	return newValue
}
