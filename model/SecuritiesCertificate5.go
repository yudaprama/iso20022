package model

// Physical representation of a security.
type SecuritiesCertificate5 struct {

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	Number *RestrictedFINXMax30Text `xml:"Nb"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	Issuer *Max4AlphaNumericText `xml:"Issr,omitempty"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (s *SecuritiesCertificate5) SetNumber(value string) {
	s.Number = (*RestrictedFINXMax30Text)(&value)
}

func (s *SecuritiesCertificate5) SetIssuer(value string) {
	s.Issuer = (*Max4AlphaNumericText)(&value)
}

func (s *SecuritiesCertificate5) SetSchemeName(value string) {
	s.SchemeName = (*Max4AlphaNumericText)(&value)
}
