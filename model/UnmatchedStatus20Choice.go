package model

// Specifies whether the status is provided with a reason or not.
type UnmatchedStatus20Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the unmatched status.
	Reason []*UnmatchedReason19 `xml:"Rsn"`
}

func (u *UnmatchedStatus20Choice) SetNoSpecifiedReason(value string) {
	u.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (u *UnmatchedStatus20Choice) AddReason() *UnmatchedReason19 {
	newValue := new(UnmatchedReason19)
	u.Reason = append(u.Reason, newValue)
	return newValue
}
