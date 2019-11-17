package model

// Specifies a shipment schedule, that is the quantity that must be shipped no sooner than the earliest shipment date and no later than the latest shipment date.
type ShipmentDateRange2 struct {

	// Sub quantity that must be shipped no sooner than the earliest shipment date and no later than the latest shipment date.
	SubQuantityValue *DecimalNumber `xml:"SubQtyVal"`

	// Earliest date whereby the goods must be shipped.
	EarliestShipmentDate *ISODate `xml:"EarlstShipmntDt,omitempty"`

	// Latest date whereby the goods must be shipped.
	LatestShipmentDate *ISODate `xml:"LatstShipmntDt,omitempty"`
}

func (s *ShipmentDateRange2) SetSubQuantityValue(value string) {
	s.SubQuantityValue = (*DecimalNumber)(&value)
}

func (s *ShipmentDateRange2) SetEarliestShipmentDate(value string) {
	s.EarliestShipmentDate = (*ISODate)(&value)
}

func (s *ShipmentDateRange2) SetLatestShipmentDate(value string) {
	s.LatestShipmentDate = (*ISODate)(&value)
}
