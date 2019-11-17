package model

// Security parameters of the ATM for the initiated key download.
type SecurityParameters4 struct {

	// Cryptographic key used to protect the key download.
	Key *CryptographicKey8 `xml:"Key,omitempty"`

	// Digital signature of implicit data depending on the security scheme download procedure.
	DigitalSignature *ContentInformationType14 `xml:"DgtlSgntr,omitempty"`

	// Ordered certificate chain of the asymmetric key encryption key, starting with the ATM certificate.
	Certificate []*Max5000Binary `xml:"Cert,omitempty"`

	// Random value from the ATM to avoid message replay.
	ATMChallenge *Max140Binary `xml:"ATMChllng,omitempty"`

	// Requested  key for downloading, depending on the key hierarchy used by the host.
	RequestedKey *Max35Text `xml:"ReqdKey,omitempty"`
}

func (s *SecurityParameters4) AddKey() *CryptographicKey8 {
	s.Key = new(CryptographicKey8)
	return s.Key
}

func (s *SecurityParameters4) AddDigitalSignature() *ContentInformationType14 {
	s.DigitalSignature = new(ContentInformationType14)
	return s.DigitalSignature
}

func (s *SecurityParameters4) AddCertificate(value string) {
	s.Certificate = append(s.Certificate, (*Max5000Binary)(&value))
}

func (s *SecurityParameters4) SetATMChallenge(value string) {
	s.ATMChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters4) SetRequestedKey(value string) {
	s.RequestedKey = (*Max35Text)(&value)
}
