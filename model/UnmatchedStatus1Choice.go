package model

// Specifies whether the status is provided with a reason or not.
type UnmatchedStatus1Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the UnmatchedStatus.
	Reason []*UnmatchedReason1 `xml:"Rsn,omitempty"`
}

func (u *UnmatchedStatus1Choice) SetNoSpecifiedReason(value string) {
	u.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (u *UnmatchedStatus1Choice) AddReason() *UnmatchedReason1 {
	newValue := new(UnmatchedReason1)
	u.Reason = append(u.Reason, newValue)
	return newValue
}
