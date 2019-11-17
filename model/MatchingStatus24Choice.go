package model

// Choice of format for the matching status.
type MatchingStatus24Choice struct {

	// Status is matched.
	Matched *ProprietaryReason4 `xml:"Mtchd"`

	// Status is unmatched.
	Unmatched *UnmatchedStatus16Choice `xml:"Umtchd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`
}

func (m *MatchingStatus24Choice) AddMatched() *ProprietaryReason4 {
	m.Matched = new(ProprietaryReason4)
	return m.Matched
}

func (m *MatchingStatus24Choice) AddUnmatched() *UnmatchedStatus16Choice {
	m.Unmatched = new(UnmatchedStatus16Choice)
	return m.Unmatched
}

func (m *MatchingStatus24Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	m.Proprietary = new(ProprietaryStatusAndReason6)
	return m.Proprietary
}
