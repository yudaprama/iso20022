package model

// Alternate identification for a party using an identification type, a country code and a text field.
type AlternatePartyIdentification5 struct {

	// Specifies the type of alternate identification of the party identified.
	IdentificationType *IdentificationType40Choice `xml:"IdTp"`

	// Nation with its own government, occupying a particular territory.
	Country *CountryCode `xml:"Ctry"`

	// Alternate identification for a party.
	AlternateIdentification *Max35Text `xml:"AltrnId"`
}

func (a *AlternatePartyIdentification5) AddIdentificationType() *IdentificationType40Choice {
	a.IdentificationType = new(IdentificationType40Choice)
	return a.IdentificationType
}

func (a *AlternatePartyIdentification5) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AlternatePartyIdentification5) SetAlternateIdentification(value string) {
	a.AlternateIdentification = (*Max35Text)(&value)
}
