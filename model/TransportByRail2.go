package model

// Information related to the transportation of goods by rail.
type TransportByRail2 struct {

	// Identifies the location where the goods are received for transportation.
	PlaceOfReceipt *Max35Text `xml:"PlcOfRct"`

	// Identifies the location of delivery of the goods.
	PlaceOfDelivery *Max35Text `xml:"PlcOfDlvry"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	RailCarrierName *Max35Text `xml:"RailCrrierNm,omitempty"`
}

func (t *TransportByRail2) SetPlaceOfReceipt(value string) {
	t.PlaceOfReceipt = (*Max35Text)(&value)
}

func (t *TransportByRail2) SetPlaceOfDelivery(value string) {
	t.PlaceOfDelivery = (*Max35Text)(&value)
}

func (t *TransportByRail2) SetRailCarrierName(value string) {
	t.RailCarrierName = (*Max35Text)(&value)
}
