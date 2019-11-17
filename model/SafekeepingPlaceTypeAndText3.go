package model

// Identification of the place of safekeeping expressed as a code and a narrative description.
type SafekeepingPlaceTypeAndText3 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace3Code `xml:"SfkpgPlcTp"`

	// Additional information about the place of safekeeping.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (s *SafekeepingPlaceTypeAndText3) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace3Code)(&value)
}

func (s *SafekeepingPlaceTypeAndText3) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}
