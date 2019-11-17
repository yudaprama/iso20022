package model

// Identification of the place of safekeeping expressed as a code and a narrative description.
type SafekeepingPlaceTypeAndText15 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace3Code `xml:"SfkpgPlcTp"`

	// Additional information about the place of safekeeping.
	Identification *RestrictedFINXMax30Text `xml:"Id,omitempty"`
}

func (s *SafekeepingPlaceTypeAndText15) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace3Code)(&value)
}

func (s *SafekeepingPlaceTypeAndText15) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax30Text)(&value)
}
