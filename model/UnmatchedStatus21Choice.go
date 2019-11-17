package model

// Specifies whether the status is provided with a reason or not.
type UnmatchedStatus21Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the unmatched status.
	Reason []*UnmatchedReason20 `xml:"Rsn"`
}

func (u *UnmatchedStatus21Choice) SetNoSpecifiedReason(value string) {
	u.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (u *UnmatchedStatus21Choice) AddReason() *UnmatchedReason20 {
	newValue := new(UnmatchedReason20)
	u.Reason = append(u.Reason, newValue)
	return newValue
}
