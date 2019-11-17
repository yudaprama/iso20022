package model

// Choice of format for the matching status.
type MatchingStatus7Choice struct {

	// Status is matched.
	Matched *ProprietaryReason1 `xml:"Mtchd"`

	// Status is unmatched.
	Unmatched *UnmatchedStatus5Choice `xml:"Umtchd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (m *MatchingStatus7Choice) AddMatched() *ProprietaryReason1 {
	m.Matched = new(ProprietaryReason1)
	return m.Matched
}

func (m *MatchingStatus7Choice) AddUnmatched() *UnmatchedStatus5Choice {
	m.Unmatched = new(UnmatchedStatus5Choice)
	return m.Unmatched
}

func (m *MatchingStatus7Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	m.Proprietary = new(ProprietaryStatusAndReason1)
	return m.Proprietary
}
