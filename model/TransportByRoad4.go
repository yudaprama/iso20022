package model

// Information related to the transportation of goods by road.
type TransportByRoad4 struct {

	// Identifies the location where the goods are received for transportation.
	PlaceOfReceipt *Max35Text `xml:"PlcOfRct"`

	// Identifies the location of delivery of the goods.
	PlaceOfDelivery *Max35Text `xml:"PlcOfDlvry"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	RoadCarrierName *Max70Text `xml:"RoadCrrierNm,omitempty"`

	// Country in which the carrier of the goods, for example, shipping company, is located or registered.
	RoadCarrierCountry *CountryCode `xml:"RoadCrrierCtry,omitempty"`

	// Name of the carrier's (for example, shipping company's) agent that acts on behalf of the carrier and may be the issuer of transport documents relating to the underlying shipment.
	CarrierAgentName *Max70Text `xml:"CrrierAgtNm,omitempty"`

	// Country of registration of the carrier's (for example, shipping company's) agent that acts on behalf of the carrier and may be the issuer of transport documents relating to the underlying shipment.
	CarrierAgentCountry *CountryCode `xml:"CrrierAgtCtry,omitempty"`
}

func (t *TransportByRoad4) SetPlaceOfReceipt(value string) {
	t.PlaceOfReceipt = (*Max35Text)(&value)
}

func (t *TransportByRoad4) SetPlaceOfDelivery(value string) {
	t.PlaceOfDelivery = (*Max35Text)(&value)
}

func (t *TransportByRoad4) SetRoadCarrierName(value string) {
	t.RoadCarrierName = (*Max70Text)(&value)
}

func (t *TransportByRoad4) SetRoadCarrierCountry(value string) {
	t.RoadCarrierCountry = (*CountryCode)(&value)
}

func (t *TransportByRoad4) SetCarrierAgentName(value string) {
	t.CarrierAgentName = (*Max70Text)(&value)
}

func (t *TransportByRoad4) SetCarrierAgentCountry(value string) {
	t.CarrierAgentCountry = (*CountryCode)(&value)
}
