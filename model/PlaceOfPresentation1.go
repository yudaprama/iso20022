package model

// Location of a presentation.
type PlaceOfPresentation1 struct {

	// Place of the presentation.
	Place *ExternalTypeOfParty1Code `xml:"Plc"`

	// Country where a presentation is to be made.
	Country *CountryCode `xml:"Ctry,omitempty"`
}

func (p *PlaceOfPresentation1) SetPlace(value string) {
	p.Place = (*ExternalTypeOfParty1Code)(&value)
}

func (p *PlaceOfPresentation1) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
