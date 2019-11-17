package model

// Choice of format for the matching status.
type MatchingStatus30Choice struct {

	// Status is matched.
	Matched *ProprietaryReason5 `xml:"Mtchd"`

	// Status is unmatched.
	Unmatched *UnmatchedStatus20Choice `xml:"Umtchd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`
}

func (m *MatchingStatus30Choice) AddMatched() *ProprietaryReason5 {
	m.Matched = new(ProprietaryReason5)
	return m.Matched
}

func (m *MatchingStatus30Choice) AddUnmatched() *UnmatchedStatus20Choice {
	m.Unmatched = new(UnmatchedStatus20Choice)
	return m.Unmatched
}

func (m *MatchingStatus30Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	m.Proprietary = new(ProprietaryStatusAndReason7)
	return m.Proprietary
}
