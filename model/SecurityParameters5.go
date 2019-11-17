package model

// Security parameters of the host downloading the key.
type SecurityParameters5 struct {

	// Random value from the host.
	HostChallenge *Max140Binary `xml:"HstChllng,omitempty"`

	// Cryptographic key used to store in the ATM.
	Key []*CryptographicKey8 `xml:"Key,omitempty"`

	// Digital signature of implicit data depending on the security scheme download procedure.
	DigitalSignature *ContentInformationType14 `xml:"DgtlSgntr,omitempty"`
}

func (s *SecurityParameters5) SetHostChallenge(value string) {
	s.HostChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters5) AddKey() *CryptographicKey8 {
	newValue := new(CryptographicKey8)
	s.Key = append(s.Key, newValue)
	return newValue
}

func (s *SecurityParameters5) AddDigitalSignature() *ContentInformationType14 {
	s.DigitalSignature = new(ContentInformationType14)
	return s.DigitalSignature
}
