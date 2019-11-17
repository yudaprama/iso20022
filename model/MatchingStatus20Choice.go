package model

// Choice of format for the matching status.
type MatchingStatus20Choice struct {

	// Status is matched.
	Matched *ProprietaryReason1 `xml:"Mtchd"`

	// Status is unmatched.
	Unmatched *UnmatchedStatus13Choice `xml:"Umtchd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (m *MatchingStatus20Choice) AddMatched() *ProprietaryReason1 {
	m.Matched = new(ProprietaryReason1)
	return m.Matched
}

func (m *MatchingStatus20Choice) AddUnmatched() *UnmatchedStatus13Choice {
	m.Unmatched = new(UnmatchedStatus13Choice)
	return m.Unmatched
}

func (m *MatchingStatus20Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	m.Proprietary = new(ProprietaryStatusAndReason1)
	return m.Proprietary
}
