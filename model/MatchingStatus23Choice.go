package model

// Choice of status for the matching.
type MatchingStatus23Choice struct {

	// Trade is matched.
	Matched *ProprietaryReason1 `xml:"Mtchd"`

	// Trade is matched with tolerance.
	MatchedWithTolerance *ProprietaryReason1 `xml:"MtchdWthTlrnce"`

	// Trade is matched alleged.
	MatchingAlleged *MatchingReason4Choice `xml:"MtchgAllgd"`

	// Trade is unmatched or mismatched.
	Unmatched *MatchingReason1Choice `xml:"Umtchd"`

	// Provides a proprietary status and a proprietary reason of the processing status of the trade.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts,omitempty"`
}

func (m *MatchingStatus23Choice) AddMatched() *ProprietaryReason1 {
	m.Matched = new(ProprietaryReason1)
	return m.Matched
}

func (m *MatchingStatus23Choice) AddMatchedWithTolerance() *ProprietaryReason1 {
	m.MatchedWithTolerance = new(ProprietaryReason1)
	return m.MatchedWithTolerance
}

func (m *MatchingStatus23Choice) AddMatchingAlleged() *MatchingReason4Choice {
	m.MatchingAlleged = new(MatchingReason4Choice)
	return m.MatchingAlleged
}

func (m *MatchingStatus23Choice) AddUnmatched() *MatchingReason1Choice {
	m.Unmatched = new(MatchingReason1Choice)
	return m.Unmatched
}

func (m *MatchingStatus23Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	m.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return m.ProprietaryStatus
}
