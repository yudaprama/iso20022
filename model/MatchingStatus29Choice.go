package model

// Choice of format for the matching status.
type MatchingStatus29Choice struct {

	// Status is matched.
	Matched *ProprietaryReason5 `xml:"Mtchd"`

	// Status is unmatched.
	Unmatched *UnmatchedStatus19Choice `xml:"Umtchd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`
}

func (m *MatchingStatus29Choice) AddMatched() *ProprietaryReason5 {
	m.Matched = new(ProprietaryReason5)
	return m.Matched
}

func (m *MatchingStatus29Choice) AddUnmatched() *UnmatchedStatus19Choice {
	m.Unmatched = new(UnmatchedStatus19Choice)
	return m.Unmatched
}

func (m *MatchingStatus29Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	m.Proprietary = new(ProprietaryStatusAndReason7)
	return m.Proprietary
}
