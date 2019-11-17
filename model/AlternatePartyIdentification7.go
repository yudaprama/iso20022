package model

// Alternate identification for a party using an identification type, a country code and a text field.
type AlternatePartyIdentification7 struct {

	// Specifies the type of alternate identification of the party identified.
	IdentificationType *IdentificationType42Choice `xml:"IdTp"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	Country *CountryCode `xml:"Ctry"`

	// Alternate identification for a party.
	AlternateIdentification *Max35Text `xml:"AltrnId"`
}

func (a *AlternatePartyIdentification7) AddIdentificationType() *IdentificationType42Choice {
	a.IdentificationType = new(IdentificationType42Choice)
	return a.IdentificationType
}

func (a *AlternatePartyIdentification7) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AlternatePartyIdentification7) SetAlternateIdentification(value string) {
	a.AlternateIdentification = (*Max35Text)(&value)
}
