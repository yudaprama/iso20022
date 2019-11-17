package model

// Information related to the transportation of goods by air.
type TransportByAir2 struct {

	// Place from where the goods must leave.
	DepartureAirport *AirportName1Choice `xml:"DprtureAirprt"`

	// Place where the goods must arrive.
	DestinationAirport *AirportName1Choice `xml:"DstnAirprt"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	AirCarrierName *Max35Text `xml:"AirCrrierNm,omitempty"`
}

func (t *TransportByAir2) AddDepartureAirport() *AirportName1Choice {
	t.DepartureAirport = new(AirportName1Choice)
	return t.DepartureAirport
}

func (t *TransportByAir2) AddDestinationAirport() *AirportName1Choice {
	t.DestinationAirport = new(AirportName1Choice)
	return t.DestinationAirport
}

func (t *TransportByAir2) SetAirCarrierName(value string) {
	t.AirCarrierName = (*Max35Text)(&value)
}
