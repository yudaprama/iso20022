package model

// Choice of format for the matching status.
type MatchingStatus26Choice struct {

	// Status is matched.
	Matched *ProprietaryReason4 `xml:"Mtchd"`

	// Status is unmatched.
	Unmatched *UnmatchedStatus18Choice `xml:"Umtchd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`
}

func (m *MatchingStatus26Choice) AddMatched() *ProprietaryReason4 {
	m.Matched = new(ProprietaryReason4)
	return m.Matched
}

func (m *MatchingStatus26Choice) AddUnmatched() *UnmatchedStatus18Choice {
	m.Unmatched = new(UnmatchedStatus18Choice)
	return m.Unmatched
}

func (m *MatchingStatus26Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	m.Proprietary = new(ProprietaryStatusAndReason6)
	return m.Proprietary
}
