package model

// Specifies individually each leg of a transport of goods.
type SingleTransport3 struct {

	// Information related to the transportation of goods by air.
	TransportByAir *TransportByAir2 `xml:"TrnsprtByAir,omitempty"`

	// Information related for the transportation of goods by sea.
	TransportBySea *TransportBySea4 `xml:"TrnsprtBySea,omitempty"`

	// Information related to the transportation of goods by road.
	TransportByRoad *TransportByRoad2 `xml:"TrnsprtByRoad,omitempty"`

	// Information related to the transportation of goods by rail.
	TransportByRail *TransportByRail2 `xml:"TrnsprtByRail,omitempty"`
}

func (s *SingleTransport3) AddTransportByAir() *TransportByAir2 {
	s.TransportByAir = new(TransportByAir2)
	return s.TransportByAir
}

func (s *SingleTransport3) AddTransportBySea() *TransportBySea4 {
	s.TransportBySea = new(TransportBySea4)
	return s.TransportBySea
}

func (s *SingleTransport3) AddTransportByRoad() *TransportByRoad2 {
	s.TransportByRoad = new(TransportByRoad2)
	return s.TransportByRoad
}

func (s *SingleTransport3) AddTransportByRail() *TransportByRail2 {
	s.TransportByRail = new(TransportByRail2)
	return s.TransportByRail
}
