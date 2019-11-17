package model

// Information related to the transportation of goods by air.
type TransportByAir4 struct {

	// Place from where the goods must leave.
	DepartureAirport *AirportName1Choice `xml:"DprtureAirprt"`

	// Place where the goods must arrive.
	DestinationAirport *AirportName1Choice `xml:"DstnAirprt"`

	// Flight number allocated by the airline that is carrying the goods from an airport of departure to an airport of destination;.
	FlightNumber *Max35Text `xml:"FlghtNb,omitempty"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	AirCarrierName *Max70Text `xml:"AirCrrierNm,omitempty"`

	// Country in which the carrier of the goods, for example, shipping company, is located or registered.
	AirCarrierCountry *CountryCode `xml:"AirCrrierCtry,omitempty"`

	// Name of the carrier's (for example, shipping company's) agent that acts on behalf of the carrier and may be the issuer of transport documents relating to the underlying shipment.
	CarrierAgentName *Max70Text `xml:"CrrierAgtNm,omitempty"`

	// Country of registration of the carrier's (for example, shipping company's) agent that acts on behalf of the carrier and may be the issuer of transport documents relating to the underlying shipment.
	CarrierAgentCountry *CountryCode `xml:"CrrierAgtCtry,omitempty"`
}

func (t *TransportByAir4) AddDepartureAirport() *AirportName1Choice {
	t.DepartureAirport = new(AirportName1Choice)
	return t.DepartureAirport
}

func (t *TransportByAir4) AddDestinationAirport() *AirportName1Choice {
	t.DestinationAirport = new(AirportName1Choice)
	return t.DestinationAirport
}

func (t *TransportByAir4) SetFlightNumber(value string) {
	t.FlightNumber = (*Max35Text)(&value)
}

func (t *TransportByAir4) SetAirCarrierName(value string) {
	t.AirCarrierName = (*Max70Text)(&value)
}

func (t *TransportByAir4) SetAirCarrierCountry(value string) {
	t.AirCarrierCountry = (*CountryCode)(&value)
}

func (t *TransportByAir4) SetCarrierAgentName(value string) {
	t.CarrierAgentName = (*Max70Text)(&value)
}

func (t *TransportByAir4) SetCarrierAgentCountry(value string) {
	t.CarrierAgentCountry = (*CountryCode)(&value)
}
