package model

// Parameters related to the security of software application and application protocol.
type SecurityParameters3 struct {

	// Version of the security parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Key to inject in the point of interaction, protected by the temporary key previously sent.
	SymmetricKey []*CryptographicKey5 `xml:"SmmtrcKey,omitempty"`
}

func (s *SecurityParameters3) SetVersion(value string) {
	s.Version = (*Max256Text)(&value)
}

func (s *SecurityParameters3) SetPOIChallenge(value string) {
	s.POIChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters3) SetTMChallenge(value string) {
	s.TMChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters3) AddSymmetricKey() *CryptographicKey5 {
	newValue := new(CryptographicKey5)
	s.SymmetricKey = append(s.SymmetricKey, newValue)
	return newValue
}
