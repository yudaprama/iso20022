package model

// Information related to the transportation of goods by air.
type TransportByAir3 struct {

	// Place from where the goods must leave.
	DepartureAirport []*AirportName1Choice `xml:"DprtureAirprt,omitempty"`

	// Place where the goods must arrive.
	DestinationAirport []*AirportName1Choice `xml:"DstnAirprt"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	AirCarrierName *Max35Text `xml:"AirCrrierNm,omitempty"`
}

func (t *TransportByAir3) AddDepartureAirport() *AirportName1Choice {
	newValue := new(AirportName1Choice)
	t.DepartureAirport = append(t.DepartureAirport, newValue)
	return newValue
}

func (t *TransportByAir3) AddDestinationAirport() *AirportName1Choice {
	newValue := new(AirportName1Choice)
	t.DestinationAirport = append(t.DestinationAirport, newValue)
	return newValue
}

func (t *TransportByAir3) SetAirCarrierName(value string) {
	t.AirCarrierName = (*Max35Text)(&value)
}
