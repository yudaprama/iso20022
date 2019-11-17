package model

// Choice of failing status.
type FailingStatus12Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the failing status.
	Reason []*FailingReason10 `xml:"Rsn"`
}

func (f *FailingStatus12Choice) SetNoSpecifiedReason(value string) {
	f.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (f *FailingStatus12Choice) AddReason() *FailingReason10 {
	newValue := new(FailingReason10)
	f.Reason = append(f.Reason, newValue)
	return newValue
}
