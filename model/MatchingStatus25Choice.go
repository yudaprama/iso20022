package model

// Choice of format for the matching status.
type MatchingStatus25Choice struct {

	// Status is matched.
	Matched *ProprietaryReason4 `xml:"Mtchd"`

	// Status is unmatched.
	Unmatched *UnmatchedStatus17Choice `xml:"Umtchd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`
}

func (m *MatchingStatus25Choice) AddMatched() *ProprietaryReason4 {
	m.Matched = new(ProprietaryReason4)
	return m.Matched
}

func (m *MatchingStatus25Choice) AddUnmatched() *UnmatchedStatus17Choice {
	m.Unmatched = new(UnmatchedStatus17Choice)
	return m.Unmatched
}

func (m *MatchingStatus25Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	m.Proprietary = new(ProprietaryStatusAndReason6)
	return m.Proprietary
}
