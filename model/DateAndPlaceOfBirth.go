package model

// Date and place of birth of a person.
type DateAndPlaceOfBirth struct {

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt"`

	// Province where a person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where a person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth"`

	// Country where a person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth"`
}

func (d *DateAndPlaceOfBirth) SetBirthDate(value string) {
	d.BirthDate = (*ISODate)(&value)
}

func (d *DateAndPlaceOfBirth) SetProvinceOfBirth(value string) {
	d.ProvinceOfBirth = (*Max35Text)(&value)
}

func (d *DateAndPlaceOfBirth) SetCityOfBirth(value string) {
	d.CityOfBirth = (*Max35Text)(&value)
}

func (d *DateAndPlaceOfBirth) SetCountryOfBirth(value string) {
	d.CountryOfBirth = (*CountryCode)(&value)
}
