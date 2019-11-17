package model

// Alternate identification for a party using an identification type, a country code and a text field.
type AlternatePartyIdentification8 struct {

	// Specifies the type of alternate identification of the party identified.
	IdentificationType *IdentificationType43Choice `xml:"IdTp"`

	// Nation with its own government, occupying a particular territory.
	Country *CountryCode `xml:"Ctry"`

	// Alternate identification for a party.
	AlternateIdentification *Max35Text `xml:"AltrnId"`
}

func (a *AlternatePartyIdentification8) AddIdentificationType() *IdentificationType43Choice {
	a.IdentificationType = new(IdentificationType43Choice)
	return a.IdentificationType
}

func (a *AlternatePartyIdentification8) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AlternatePartyIdentification8) SetAlternateIdentification(value string) {
	a.AlternateIdentification = (*Max35Text)(&value)
}
