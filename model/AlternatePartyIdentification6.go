package model

// Alternate identification for a party using an id type, a country code and an text field.
type AlternatePartyIdentification6 struct {

	// Specifies the type of alternate identification which can be used to give an alternate identification of the party identified.
	TypeOfIdentification *IdentificationType41Choice `xml:"TpOfId"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	Country *CountryCode `xml:"Ctry"`

	// Alternate identification for a party.
	AlternateIdentification *Max35Text `xml:"AltrnId"`
}

func (a *AlternatePartyIdentification6) AddTypeOfIdentification() *IdentificationType41Choice {
	a.TypeOfIdentification = new(IdentificationType41Choice)
	return a.TypeOfIdentification
}

func (a *AlternatePartyIdentification6) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AlternatePartyIdentification6) SetAlternateIdentification(value string) {
	a.AlternateIdentification = (*Max35Text)(&value)
}
