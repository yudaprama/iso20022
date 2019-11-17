package model

// Choice of format for the matching status.
type MatchingStatus12Choice struct {

	// Status is matched.
	Matched *ProprietaryReason1 `xml:"Mtchd"`

	// Status is unmatched.
	Unmatched *UnmatchedStatus6Choice `xml:"Umtchd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (m *MatchingStatus12Choice) AddMatched() *ProprietaryReason1 {
	m.Matched = new(ProprietaryReason1)
	return m.Matched
}

func (m *MatchingStatus12Choice) AddUnmatched() *UnmatchedStatus6Choice {
	m.Unmatched = new(UnmatchedStatus6Choice)
	return m.Unmatched
}

func (m *MatchingStatus12Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	m.Proprietary = new(ProprietaryStatusAndReason1)
	return m.Proprietary
}
