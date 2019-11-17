package model

// Specifies an earliest shipment date and a latest shipment date.
type ShipmentDateRange1 struct {

	// Earliest date whereby the goods must be shipped.
	EarliestShipmentDate *ISODate `xml:"EarlstShipmntDt,omitempty"`

	// Latest date whereby the goods must be shipped.
	LatestShipmentDate *ISODate `xml:"LatstShipmntDt,omitempty"`
}

func (s *ShipmentDateRange1) SetEarliestShipmentDate(value string) {
	s.EarliestShipmentDate = (*ISODate)(&value)
}

func (s *ShipmentDateRange1) SetLatestShipmentDate(value string) {
	s.LatestShipmentDate = (*ISODate)(&value)
}
