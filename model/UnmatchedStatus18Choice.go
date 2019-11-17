package model

// Specifies whether the status is provided with a reason or not.
type UnmatchedStatus18Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the unmatched status.
	Reason []*UnmatchedReason17 `xml:"Rsn"`
}

func (u *UnmatchedStatus18Choice) SetNoSpecifiedReason(value string) {
	u.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (u *UnmatchedStatus18Choice) AddReason() *UnmatchedReason17 {
	newValue := new(UnmatchedReason17)
	u.Reason = append(u.Reason, newValue)
	return newValue
}
