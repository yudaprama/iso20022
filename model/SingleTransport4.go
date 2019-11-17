package model

// Specifies individually each leg of a transport of goods.
type SingleTransport4 struct {

	// Moving of goods or people from one place to another by vehicle.
	TransportByAir []*TransportByAir3 `xml:"TrnsprtByAir,omitempty"`

	// Information related for the transportation of goods by sea.
	TransportBySea []*TransportBySea3 `xml:"TrnsprtBySea,omitempty"`

	// Information related to the transportation of goods by road.
	TransportByRoad []*TransportByRoad3 `xml:"TrnsprtByRoad,omitempty"`

	// Information related to the transportation of goods by rail.
	TransportByRail []*TransportByRail3 `xml:"TrnsprtByRail,omitempty"`
}

func (s *SingleTransport4) AddTransportByAir() *TransportByAir3 {
	newValue := new(TransportByAir3)
	s.TransportByAir = append(s.TransportByAir, newValue)
	return newValue
}

func (s *SingleTransport4) AddTransportBySea() *TransportBySea3 {
	newValue := new(TransportBySea3)
	s.TransportBySea = append(s.TransportBySea, newValue)
	return newValue
}

func (s *SingleTransport4) AddTransportByRoad() *TransportByRoad3 {
	newValue := new(TransportByRoad3)
	s.TransportByRoad = append(s.TransportByRoad, newValue)
	return newValue
}

func (s *SingleTransport4) AddTransportByRail() *TransportByRail3 {
	newValue := new(TransportByRail3)
	s.TransportByRail = append(s.TransportByRail, newValue)
	return newValue
}
