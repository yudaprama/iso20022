package model

// Specifies whether the status is provided with a reason or not.
type UnmatchedStatus2Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the UnmatchedStatus.
	Reason []*UnmatchedReason2 `xml:"Rsn,omitempty"`
}

func (u *UnmatchedStatus2Choice) SetNoSpecifiedReason(value string) {
	u.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (u *UnmatchedStatus2Choice) AddReason() *UnmatchedReason2 {
	newValue := new(UnmatchedReason2)
	u.Reason = append(u.Reason, newValue)
	return newValue
}
