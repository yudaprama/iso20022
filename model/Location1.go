package model

// Location information.
type Location1 struct {

	// Country of jurisdiction.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Codified representation of the jurisdiction as published in ISO 3166-2.
	CountrySubDivision *CountrySubdivision1Choice `xml:"CtrySubDvsn,omitempty"`

	// Name of jurisdiction, for example, Frankfurt.
	Text []*Max2000Text `xml:"Txt,omitempty"`
}

func (l *Location1) SetCountry(value string) {
	l.Country = (*CountryCode)(&value)
}

func (l *Location1) AddCountrySubDivision() *CountrySubdivision1Choice {
	l.CountrySubDivision = new(CountrySubdivision1Choice)
	return l.CountrySubDivision
}

func (l *Location1) AddText(value string) {
	l.Text = append(l.Text, (*Max2000Text)(&value))
}
