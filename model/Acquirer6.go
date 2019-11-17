package model

// Acquirer of the card transaction.
type Acquirer6 struct {

	// Identification of the acquirer.
	// It correspond to the ISO 8583 field number 32.
	Identification *Max35Text `xml:"Id"`

	// Identification of the entity assigning the acquirer identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Country of the acquirer.
	// It correspond to the ISO 8583 field number 19.
	CountryCode *ISO3NumericCountryCode `xml:"CtryCd,omitempty"`
}

func (a *Acquirer6) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Acquirer6) SetIssuer(value string) {
	a.Issuer = (*Max35Text)(&value)
}

func (a *Acquirer6) SetCountryCode(value string) {
	a.CountryCode = (*ISO3NumericCountryCode)(&value)
}
