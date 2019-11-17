package model

// Alternate identification for a party using an identification type, a country code and a text field.
type AlternatePartyIdentification9 struct {

	// Specifies the type of alternate identification of the party identified.
	IdentificationType *IdentificationType44Choice `xml:"IdTp"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	Country *CountryCode `xml:"Ctry"`

	// Alternate identification for a party.
	AlternateIdentification *RestrictedFINXMax30Text `xml:"AltrnId"`
}

func (a *AlternatePartyIdentification9) AddIdentificationType() *IdentificationType44Choice {
	a.IdentificationType = new(IdentificationType44Choice)
	return a.IdentificationType
}

func (a *AlternatePartyIdentification9) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AlternatePartyIdentification9) SetAlternateIdentification(value string) {
	a.AlternateIdentification = (*RestrictedFINXMax30Text)(&value)
}
