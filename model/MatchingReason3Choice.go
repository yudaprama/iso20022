package model

// Choice of format for the matching processing status.
type MatchingReason3Choice struct {

	// Specifies the reason of the MatchedAlleged Status.
	Reason []*AllegmentMatchingReason1 `xml:"Rsn"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (m *MatchingReason3Choice) AddReason() *AllegmentMatchingReason1 {
	newValue := new(AllegmentMatchingReason1)
	m.Reason = append(m.Reason, newValue)
	return newValue
}

func (m *MatchingReason3Choice) SetNoSpecifiedReason(value string) {
	m.NoSpecifiedReason = (*NoReasonCode)(&value)
}
