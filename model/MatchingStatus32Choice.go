package model

// Choice of format for the matching status.
type MatchingStatus32Choice struct {

	// Status is matched.
	Matched *ProprietaryReason5 `xml:"Mtchd"`

	// Status is unmatched.
	Unmatched *UnmatchedStatus21Choice `xml:"Umtchd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`
}

func (m *MatchingStatus32Choice) AddMatched() *ProprietaryReason5 {
	m.Matched = new(ProprietaryReason5)
	return m.Matched
}

func (m *MatchingStatus32Choice) AddUnmatched() *UnmatchedStatus21Choice {
	m.Unmatched = new(UnmatchedStatus21Choice)
	return m.Unmatched
}

func (m *MatchingStatus32Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	m.Proprietary = new(ProprietaryStatusAndReason7)
	return m.Proprietary
}
