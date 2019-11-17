package model

// Identification of the place of safekeeping expressed as a code and a BIC.
type SafekeepingPlaceAsCodeAndPartyIdentification struct {

	// Place of safekeeping as a code.
	PlaceSafekeeping *SafekeepingPlace1Code `xml:"PlcSfkpg"`

	// Additional information about the place of safekeeping.
	Narrative *Max35Text `xml:"Nrrtv,omitempty"`

	// Place of safekeeping.
	Party *PartyIdentification3 `xml:"Pty,omitempty"`
}

func (s *SafekeepingPlaceAsCodeAndPartyIdentification) SetPlaceSafekeeping(value string) {
	s.PlaceSafekeeping = (*SafekeepingPlace1Code)(&value)
}

func (s *SafekeepingPlaceAsCodeAndPartyIdentification) SetNarrative(value string) {
	s.Narrative = (*Max35Text)(&value)
}

func (s *SafekeepingPlaceAsCodeAndPartyIdentification) AddParty() *PartyIdentification3 {
	s.Party = new(PartyIdentification3)
	return s.Party
}
