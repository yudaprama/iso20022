package model

// Choice of failing status.
type FailingStatus11Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the failing status.
	Reason []*FailingReason9 `xml:"Rsn"`
}

func (f *FailingStatus11Choice) SetNoSpecifiedReason(value string) {
	f.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (f *FailingStatus11Choice) AddReason() *FailingReason9 {
	newValue := new(FailingReason9)
	f.Reason = append(f.Reason, newValue)
	return newValue
}
