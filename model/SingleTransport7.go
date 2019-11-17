package model

// Specifies individually each leg of a transport of goods.
type SingleTransport7 struct {

	// Moving of goods or people from one place to another by vehicle.
	TransportByAir []*TransportByAir5 `xml:"TrnsprtByAir,omitempty"`

	// Information related for the transportation of goods by sea.
	TransportBySea []*TransportBySea6 `xml:"TrnsprtBySea,omitempty"`

	// Information related to the transportation of goods by road.
	TransportByRoad []*TransportByRoad5 `xml:"TrnsprtByRoad,omitempty"`

	// Information related to the transportation of goods by rail.
	TransportByRail []*TransportByRail5 `xml:"TrnsprtByRail,omitempty"`
}

func (s *SingleTransport7) AddTransportByAir() *TransportByAir5 {
	newValue := new(TransportByAir5)
	s.TransportByAir = append(s.TransportByAir, newValue)
	return newValue
}

func (s *SingleTransport7) AddTransportBySea() *TransportBySea6 {
	newValue := new(TransportBySea6)
	s.TransportBySea = append(s.TransportBySea, newValue)
	return newValue
}

func (s *SingleTransport7) AddTransportByRoad() *TransportByRoad5 {
	newValue := new(TransportByRoad5)
	s.TransportByRoad = append(s.TransportByRoad, newValue)
	return newValue
}

func (s *SingleTransport7) AddTransportByRail() *TransportByRail5 {
	newValue := new(TransportByRail5)
	s.TransportByRail = append(s.TransportByRail, newValue)
	return newValue
}
