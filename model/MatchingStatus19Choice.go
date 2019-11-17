package model

// Choice of format for the matching status.
type MatchingStatus19Choice struct {

	// Status is matched.
	Matched *ProprietaryReason1 `xml:"Mtchd"`

	// Status is unmatched.
	Unmatched *UnmatchedStatus12Choice `xml:"Umtchd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (m *MatchingStatus19Choice) AddMatched() *ProprietaryReason1 {
	m.Matched = new(ProprietaryReason1)
	return m.Matched
}

func (m *MatchingStatus19Choice) AddUnmatched() *UnmatchedStatus12Choice {
	m.Unmatched = new(UnmatchedStatus12Choice)
	return m.Unmatched
}

func (m *MatchingStatus19Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	m.Proprietary = new(ProprietaryStatusAndReason1)
	return m.Proprietary
}
