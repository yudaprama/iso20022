package model

// Information related to the transportation of goods by road.
type TransportByRoad3 struct {

	// Identifies the location where the goods are received for transportation.
	PlaceOfReceipt []*Max35Text `xml:"PlcOfRct,omitempty"`

	// Identifies the location of delivery of the goods.
	PlaceOfDelivery []*Max35Text `xml:"PlcOfDlvry"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	RoadCarrierName *Max35Text `xml:"RoadCrrierNm,omitempty"`
}

func (t *TransportByRoad3) AddPlaceOfReceipt(value string) {
	t.PlaceOfReceipt = append(t.PlaceOfReceipt, (*Max35Text)(&value))
}

func (t *TransportByRoad3) AddPlaceOfDelivery(value string) {
	t.PlaceOfDelivery = append(t.PlaceOfDelivery, (*Max35Text)(&value))
}

func (t *TransportByRoad3) SetRoadCarrierName(value string) {
	t.RoadCarrierName = (*Max35Text)(&value)
}
