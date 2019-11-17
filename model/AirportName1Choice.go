package model

// Identifies an airport by its code or by its name and the town where it is located.
type AirportName1Choice struct {

	// Identifies an airport by means of its IATA identification code. Example: LHR.
	AirportCode *Max6Text `xml:"AirprtCd"`

	// Identifies an airport by its location and by its name.
	OtherAirportDescription *AirportDescription1 `xml:"OthrAirprtDesc"`
}

func (a *AirportName1Choice) SetAirportCode(value string) {
	a.AirportCode = (*Max6Text)(&value)
}

func (a *AirportName1Choice) AddOtherAirportDescription() *AirportDescription1 {
	a.OtherAirportDescription = new(AirportDescription1)
	return a.OtherAirportDescription
}
