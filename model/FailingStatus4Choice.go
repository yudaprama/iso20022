package model

// Choice of failing status.
type FailingStatus4Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the FailingStatus.
	Reason []*FailingReason4 `xml:"Rsn"`
}

func (f *FailingStatus4Choice) SetNoSpecifiedReason(value string) {
	f.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (f *FailingStatus4Choice) AddReason() *FailingReason4 {
	newValue := new(FailingReason4)
	f.Reason = append(f.Reason, newValue)
	return newValue
}
