package model

// Point of interaction parameters related to the security of software application and application protocol.
type SecurityParameters1 struct {

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Key to inject in the point of interaction, protected by the temporary key previously sent.
	SymmetricKey []*CryptographicKey2 `xml:"SmmtrcKey,omitempty"`
}

func (s *SecurityParameters1) SetPOIChallenge(value string) {
	s.POIChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters1) SetTMChallenge(value string) {
	s.TMChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters1) AddSymmetricKey() *CryptographicKey2 {
	newValue := new(CryptographicKey2)
	s.SymmetricKey = append(s.SymmetricKey, newValue)
	return newValue
}
