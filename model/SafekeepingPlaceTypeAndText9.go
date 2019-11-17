package model

// Identification of the place of safekeeping expressed as a code and a narrative description.
type SafekeepingPlaceTypeAndText9 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace2Code `xml:"SfkpgPlcTp"`

	// Additional information about the place of safekeeping.
	Identification *RestrictedFINXMax30Text `xml:"Id,omitempty"`
}

func (s *SafekeepingPlaceTypeAndText9) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace2Code)(&value)
}

func (s *SafekeepingPlaceTypeAndText9) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax30Text)(&value)
}
