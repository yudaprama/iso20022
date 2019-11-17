package model

// Choice of format for the matching status.
type MatchingStatus3Choice struct {

	// Status is matched.
	Matched *NoSpecifiedReason1 `xml:"Mtchd"`

	// Status is unmatched.
	Unmatched *UnmatchedStatus2Choice `xml:"Umtchd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (m *MatchingStatus3Choice) AddMatched() *NoSpecifiedReason1 {
	m.Matched = new(NoSpecifiedReason1)
	return m.Matched
}

func (m *MatchingStatus3Choice) AddUnmatched() *UnmatchedStatus2Choice {
	m.Unmatched = new(UnmatchedStatus2Choice)
	return m.Unmatched
}

func (m *MatchingStatus3Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	m.Proprietary = new(ProprietaryStatusAndReason1)
	return m.Proprietary
}
