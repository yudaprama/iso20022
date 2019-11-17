package model

// Identifies an airport by its location and by its name.
type AirportDescription1 struct {

	// Identifies the town where the airport is located. For example: London.
	Town *Max35Text `xml:"Twn"`

	// Identifies the airport by its name. For example: Heathrow.
	AirportName *Max35Text `xml:"AirprtNm,omitempty"`
}

func (a *AirportDescription1) SetTown(value string) {
	a.Town = (*Max35Text)(&value)
}

func (a *AirportDescription1) SetAirportName(value string) {
	a.AirportName = (*Max35Text)(&value)
}
