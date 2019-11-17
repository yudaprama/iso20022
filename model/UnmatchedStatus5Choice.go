package model

// Specifies whether the status is provided with a reason or not.
type UnmatchedStatus5Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the UnmatchedStatus.
	Reason []*UnmatchedReason6 `xml:"Rsn"`
}

func (u *UnmatchedStatus5Choice) SetNoSpecifiedReason(value string) {
	u.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (u *UnmatchedStatus5Choice) AddReason() *UnmatchedReason6 {
	newValue := new(UnmatchedReason6)
	u.Reason = append(u.Reason, newValue)
	return newValue
}
