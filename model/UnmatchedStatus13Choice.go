package model

// Specifies whether the status is provided with a reason or not.
type UnmatchedStatus13Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the UnmatchedStatus.
	Reason []*UnmatchedReason12 `xml:"Rsn"`
}

func (u *UnmatchedStatus13Choice) SetNoSpecifiedReason(value string) {
	u.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (u *UnmatchedStatus13Choice) AddReason() *UnmatchedReason12 {
	newValue := new(UnmatchedReason12)
	u.Reason = append(u.Reason, newValue)
	return newValue
}
