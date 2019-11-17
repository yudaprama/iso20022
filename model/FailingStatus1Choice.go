package model

// Choice of failing status.
type FailingStatus1Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the FailingStatus.
	Reason []*FailingReason1 `xml:"Rsn,omitempty"`
}

func (f *FailingStatus1Choice) SetNoSpecifiedReason(value string) {
	f.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (f *FailingStatus1Choice) AddReason() *FailingReason1 {
	newValue := new(FailingReason1)
	f.Reason = append(f.Reason, newValue)
	return newValue
}
