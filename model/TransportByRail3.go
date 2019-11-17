package model

// Information related to the transportation of goods by rail.
type TransportByRail3 struct {

	// Identifies the location where the goods are received for transportation.
	PlaceOfReceipt []*Max35Text `xml:"PlcOfRct,omitempty"`

	// Identifies the location of delivery of the goods.
	PlaceOfDelivery []*Max35Text `xml:"PlcOfDlvry"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	RailCarrierName *Max35Text `xml:"RailCrrierNm,omitempty"`
}

func (t *TransportByRail3) AddPlaceOfReceipt(value string) {
	t.PlaceOfReceipt = append(t.PlaceOfReceipt, (*Max35Text)(&value))
}

func (t *TransportByRail3) AddPlaceOfDelivery(value string) {
	t.PlaceOfDelivery = append(t.PlaceOfDelivery, (*Max35Text)(&value))
}

func (t *TransportByRail3) SetRailCarrierName(value string) {
	t.RailCarrierName = (*Max35Text)(&value)
}
