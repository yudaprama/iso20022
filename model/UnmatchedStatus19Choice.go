package model

// Specifies whether the status is provided with a reason or not.
type UnmatchedStatus19Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the unmatched status.
	Reason []*UnmatchedReason18 `xml:"Rsn"`
}

func (u *UnmatchedStatus19Choice) SetNoSpecifiedReason(value string) {
	u.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (u *UnmatchedStatus19Choice) AddReason() *UnmatchedReason18 {
	newValue := new(UnmatchedReason18)
	u.Reason = append(u.Reason, newValue)
	return newValue
}
