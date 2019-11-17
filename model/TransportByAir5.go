package model

// Information related to the transportation of goods by air.
type TransportByAir5 struct {

	// Place from where the goods must leave.
	DepartureAirport []*AirportName1Choice `xml:"DprtureAirprt,omitempty"`

	// Place where the goods must arrive.
	DestinationAirport []*AirportName1Choice `xml:"DstnAirprt"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	AirCarrierName *Max70Text `xml:"AirCrrierNm,omitempty"`

	// Country in which the carrier of the goods, for example, shipping company, is located or registered.
	AirCarrierCountry *CountryCode `xml:"AirCrrierCtry,omitempty"`

	// Name of the carrier's (for example, shipping company's) agent that acts on behalf of the carrier and may be the issuer of transport documents relating to the underlying shipment.
	CarrierAgentName *Max70Text `xml:"CrrierAgtNm,omitempty"`

	// Country of registration of the carrier's (for example, shipping company's) agent that acts on behalf of the carrier and may be the issuer of transport documents relating to the underlying shipment.
	CarrierAgentCountry *CountryCode `xml:"CrrierAgtCtry,omitempty"`
}

func (t *TransportByAir5) AddDepartureAirport() *AirportName1Choice {
	newValue := new(AirportName1Choice)
	t.DepartureAirport = append(t.DepartureAirport, newValue)
	return newValue
}

func (t *TransportByAir5) AddDestinationAirport() *AirportName1Choice {
	newValue := new(AirportName1Choice)
	t.DestinationAirport = append(t.DestinationAirport, newValue)
	return newValue
}

func (t *TransportByAir5) SetAirCarrierName(value string) {
	t.AirCarrierName = (*Max70Text)(&value)
}

func (t *TransportByAir5) SetAirCarrierCountry(value string) {
	t.AirCarrierCountry = (*CountryCode)(&value)
}

func (t *TransportByAir5) SetCarrierAgentName(value string) {
	t.CarrierAgentName = (*Max70Text)(&value)
}

func (t *TransportByAir5) SetCarrierAgentCountry(value string) {
	t.CarrierAgentCountry = (*CountryCode)(&value)
}
