package model

// Physical representation of a security.
type SecuritiesCertificate3 struct {

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	Number *Max35Text `xml:"Nb"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (s *SecuritiesCertificate3) SetNumber(value string) {
	s.Number = (*Max35Text)(&value)
}

func (s *SecuritiesCertificate3) SetIssuer(value string) {
	s.Issuer = (*Max35Text)(&value)
}

func (s *SecuritiesCertificate3) SetSchemeName(value string) {
	s.SchemeName = (*Max35Text)(&value)
}
