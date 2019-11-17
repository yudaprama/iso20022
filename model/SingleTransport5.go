package model

// Specifies individually each leg of a transport of goods.
type SingleTransport5 struct {

	// Information related to the transportation of goods by air.
	TransportByAir []*TransportByAir2 `xml:"TrnsprtByAir,omitempty"`

	// Information related for the transportation of goods by sea.
	TransportBySea []*TransportBySea4 `xml:"TrnsprtBySea,omitempty"`

	// Moving of goods or people from one place to another by vehicle.
	TransportByRoad []*TransportByRoad2 `xml:"TrnsprtByRoad,omitempty"`

	// Information related to the transportation of goods by rail.
	TransportByRail []*TransportByRail2 `xml:"TrnsprtByRail,omitempty"`
}

func (s *SingleTransport5) AddTransportByAir() *TransportByAir2 {
	newValue := new(TransportByAir2)
	s.TransportByAir = append(s.TransportByAir, newValue)
	return newValue
}

func (s *SingleTransport5) AddTransportBySea() *TransportBySea4 {
	newValue := new(TransportBySea4)
	s.TransportBySea = append(s.TransportBySea, newValue)
	return newValue
}

func (s *SingleTransport5) AddTransportByRoad() *TransportByRoad2 {
	newValue := new(TransportByRoad2)
	s.TransportByRoad = append(s.TransportByRoad, newValue)
	return newValue
}

func (s *SingleTransport5) AddTransportByRail() *TransportByRail2 {
	newValue := new(TransportByRail2)
	s.TransportByRail = append(s.TransportByRail, newValue)
	return newValue
}
