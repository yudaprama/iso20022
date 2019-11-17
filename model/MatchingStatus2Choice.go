package model

// Choice of format for the matching status.
type MatchingStatus2Choice struct {

	// Status is matched.
	Matched *NoSpecifiedReason1 `xml:"Mtchd"`

	// Status is unmatched.
	Unmatched *UnmatchedStatus1Choice `xml:"Umtchd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (m *MatchingStatus2Choice) AddMatched() *NoSpecifiedReason1 {
	m.Matched = new(NoSpecifiedReason1)
	return m.Matched
}

func (m *MatchingStatus2Choice) AddUnmatched() *UnmatchedStatus1Choice {
	m.Unmatched = new(UnmatchedStatus1Choice)
	return m.Unmatched
}

func (m *MatchingStatus2Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	m.Proprietary = new(ProprietaryStatusAndReason1)
	return m.Proprietary
}
