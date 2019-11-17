package model

// Specifies whether the status is provided with a reason or not.
type UnmatchedStatus17Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the unmatched status.
	Reason []*UnmatchedReason16 `xml:"Rsn"`
}

func (u *UnmatchedStatus17Choice) SetNoSpecifiedReason(value string) {
	u.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (u *UnmatchedStatus17Choice) AddReason() *UnmatchedReason16 {
	newValue := new(UnmatchedReason16)
	u.Reason = append(u.Reason, newValue)
	return newValue
}
