package model

// Information related to the transportation of goods by road.
type TransportByRoad2 struct {

	// Identifies the location where the goods are received for transportation.
	PlaceOfReceipt *Max35Text `xml:"PlcOfRct"`

	// Identifies the location of delivery of the goods.
	PlaceOfDelivery *Max35Text `xml:"PlcOfDlvry"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	RoadCarrierName *Max35Text `xml:"RoadCrrierNm,omitempty"`
}

func (t *TransportByRoad2) SetPlaceOfReceipt(value string) {
	t.PlaceOfReceipt = (*Max35Text)(&value)
}

func (t *TransportByRoad2) SetPlaceOfDelivery(value string) {
	t.PlaceOfDelivery = (*Max35Text)(&value)
}

func (t *TransportByRoad2) SetRoadCarrierName(value string) {
	t.RoadCarrierName = (*Max35Text)(&value)
}
