package model

// Point of interaction parameters related to the security of software application and application protocol.
type SecurityParameters2 struct {

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Key to inject in the point of interaction, protected by the temporary key previously sent.
	SymmetricKey []*CryptographicKey4 `xml:"SmmtrcKey,omitempty"`
}

func (s *SecurityParameters2) SetPOIChallenge(value string) {
	s.POIChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters2) SetTMChallenge(value string) {
	s.TMChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters2) AddSymmetricKey() *CryptographicKey4 {
	newValue := new(CryptographicKey4)
	s.SymmetricKey = append(s.SymmetricKey, newValue)
	return newValue
}
