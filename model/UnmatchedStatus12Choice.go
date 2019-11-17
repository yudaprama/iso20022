package model

// Specifies whether the status is provided with a reason or not.
type UnmatchedStatus12Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the UnmatchedStatus.
	Reason []*UnmatchedReason11 `xml:"Rsn"`
}

func (u *UnmatchedStatus12Choice) SetNoSpecifiedReason(value string) {
	u.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (u *UnmatchedStatus12Choice) AddReason() *UnmatchedReason11 {
	newValue := new(UnmatchedReason11)
	u.Reason = append(u.Reason, newValue)
	return newValue
}
