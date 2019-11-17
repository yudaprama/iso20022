package model

// Defines an identifier for a party relative to another party using an identifier of another party followed by a local identifier issued by the other party.
// It is assumed that customers of an identifiable party can be referenced by an identifier relative to that party. The identification of the party together with the relative identifier identifies the customer.
// Such references can occur in sequence. The last occurrence of RelativeIdentifier is the local identifier at the party recursively defined by the PrefixParty and all preceding occurrences of RelativeIdentifier.
// Technical note: The sequence of relative identifiers is used to avoid a recursive definition in the catalogue.
type SingleQualifiedPartyIdentification1 struct {

	// Party identification recognisable by parties in the trade.
	BaseParty *TradeParty1 `xml:"BasePty"`

	// Identifies a party, each identifier is recursively defined relative to the party identified by the base party and the preceding part of the list.
	RelativeIdentifier []*Max35Text `xml:"RltvIdr,omitempty"`
}

func (s *SingleQualifiedPartyIdentification1) AddBaseParty() *TradeParty1 {
	s.BaseParty = new(TradeParty1)
	return s.BaseParty
}

func (s *SingleQualifiedPartyIdentification1) AddRelativeIdentifier(value string) {
	s.RelativeIdentifier = append(s.RelativeIdentifier, (*Max35Text)(&value))
}
