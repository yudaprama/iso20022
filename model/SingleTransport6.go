package model

// Specifies individually each leg of a transport of goods.
type SingleTransport6 struct {

	// Information related to the transportation of goods by air.
	TransportByAir []*TransportByAir4 `xml:"TrnsprtByAir,omitempty"`

	// Information related for the transportation of goods by sea.
	TransportBySea []*TransportBySea5 `xml:"TrnsprtBySea,omitempty"`

	// Moving of goods or people from one place to another by vehicle.
	TransportByRoad []*TransportByRoad4 `xml:"TrnsprtByRoad,omitempty"`

	// Information related to the transportation of goods by rail.
	TransportByRail []*TransportByRail4 `xml:"TrnsprtByRail,omitempty"`
}

func (s *SingleTransport6) AddTransportByAir() *TransportByAir4 {
	newValue := new(TransportByAir4)
	s.TransportByAir = append(s.TransportByAir, newValue)
	return newValue
}

func (s *SingleTransport6) AddTransportBySea() *TransportBySea5 {
	newValue := new(TransportBySea5)
	s.TransportBySea = append(s.TransportBySea, newValue)
	return newValue
}

func (s *SingleTransport6) AddTransportByRoad() *TransportByRoad4 {
	newValue := new(TransportByRoad4)
	s.TransportByRoad = append(s.TransportByRoad, newValue)
	return newValue
}

func (s *SingleTransport6) AddTransportByRail() *TransportByRail4 {
	newValue := new(TransportByRail4)
	s.TransportByRail = append(s.TransportByRail, newValue)
	return newValue
}
