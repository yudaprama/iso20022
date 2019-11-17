package model

// Identification of the place of safekeeping expressed as a code and a narrative description.
type SafekeepingPlaceTypeAndText6 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace2Code `xml:"SfkpgPlcTp"`

	// Additional information about the place of safekeeping.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (s *SafekeepingPlaceTypeAndText6) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace2Code)(&value)
}

func (s *SafekeepingPlaceTypeAndText6) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}
