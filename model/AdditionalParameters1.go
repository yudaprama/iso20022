package model

// Sort criteria.
type AdditionalParameters1 struct {

	// Specifies the country.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Specifies the currency.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Specifies the geographical area, eg, Asia-Pacific, Europe, Middle-East.
	GeographicalArea *Max35Text `xml:"GeoArea,omitempty"`
}

func (a *AdditionalParameters1) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AdditionalParameters1) SetCurrency(value string) {
	a.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (a *AdditionalParameters1) SetGeographicalArea(value string) {
	a.GeographicalArea = (*Max35Text)(&value)
}
