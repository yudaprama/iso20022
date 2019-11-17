package model

// Place identification of the place of safekeeping expressed as a code and a BIC.
type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace1Code `xml:"SfkpgPlcTp"`

	// Place of safekeeping.
	Identification *AnyBICIdentifier `xml:"Id"`
}

func (s *SafekeepingPlaceTypeAndAnyBICIdentifier1) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace1Code)(&value)
}

func (s *SafekeepingPlaceTypeAndAnyBICIdentifier1) SetIdentification(value string) {
	s.Identification = (*AnyBICIdentifier)(&value)
}
