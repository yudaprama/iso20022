package model

// Tax identification information.
type TaxIdentification2 struct {

	// Tax identification number or identifier.
	Identification *Max35Text `xml:"Id"`

	// Type of tax identification number or identifier.
	TaxIdentificationType *TaxIdentificationType1Choice `xml:"TaxIdTp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Date at which the identification was issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date at which the identification expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Country that issued the tax identification.
	IssuerCountry *CountryCode `xml:"IssrCtry"`
}

func (t *TaxIdentification2) SetIdentification(value string) {
	t.Identification = (*Max35Text)(&value)
}

func (t *TaxIdentification2) AddTaxIdentificationType() *TaxIdentificationType1Choice {
	t.TaxIdentificationType = new(TaxIdentificationType1Choice)
	return t.TaxIdentificationType
}

func (t *TaxIdentification2) SetIssuer(value string) {
	t.Issuer = (*Max35Text)(&value)
}

func (t *TaxIdentification2) SetIssueDate(value string) {
	t.IssueDate = (*ISODate)(&value)
}

func (t *TaxIdentification2) SetExpiryDate(value string) {
	t.ExpiryDate = (*ISODate)(&value)
}

func (t *TaxIdentification2) SetIssuerCountry(value string) {
	t.IssuerCountry = (*CountryCode)(&value)
}
