package model

// Choice of format for the matching processing status.
type MatchingReason4Choice struct {

	// Specifies the reason of the MatchedAlleged Status.
	Reason []*AllegementMatchingReason1 `xml:"Rsn"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (m *MatchingReason4Choice) AddReason() *AllegementMatchingReason1 {
	newValue := new(AllegementMatchingReason1)
	m.Reason = append(m.Reason, newValue)
	return newValue
}

func (m *MatchingReason4Choice) SetNoSpecifiedReason(value string) {
	m.NoSpecifiedReason = (*NoReasonCode)(&value)
}
