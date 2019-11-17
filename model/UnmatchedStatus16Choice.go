package model

// Specifies whether the status is provided with a reason or not.
type UnmatchedStatus16Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the unmatched status.
	Reason []*UnmatchedReason15 `xml:"Rsn"`
}

func (u *UnmatchedStatus16Choice) SetNoSpecifiedReason(value string) {
	u.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (u *UnmatchedStatus16Choice) AddReason() *UnmatchedReason15 {
	newValue := new(UnmatchedReason15)
	u.Reason = append(u.Reason, newValue)
	return newValue
}
