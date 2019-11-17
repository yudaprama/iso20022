package model

// Specifies whether the status is provided with a reason or not.
type UnmatchedStatus6Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the UnmatchedStatus.
	Reason []*UnmatchedReason7 `xml:"Rsn"`
}

func (u *UnmatchedStatus6Choice) SetNoSpecifiedReason(value string) {
	u.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (u *UnmatchedStatus6Choice) AddReason() *UnmatchedReason7 {
	newValue := new(UnmatchedReason7)
	u.Reason = append(u.Reason, newValue)
	return newValue
}
