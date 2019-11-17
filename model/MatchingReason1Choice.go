package model

// Choice of format for the matching processing status.
type MatchingReason1Choice struct {

	// Specifies the reason of the UnmatchStatus.
	Reason []*UnmatchedReason5 `xml:"Rsn"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (m *MatchingReason1Choice) AddReason() *UnmatchedReason5 {
	newValue := new(UnmatchedReason5)
	m.Reason = append(m.Reason, newValue)
	return newValue
}

func (m *MatchingReason1Choice) SetNoSpecifiedReason(value string) {
	m.NoSpecifiedReason = (*NoReasonCode)(&value)
}
