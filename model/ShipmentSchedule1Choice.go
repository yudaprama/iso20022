package model

// Choice between earliest/latest shipment date and a shipment schedule per sub quantity of line item quantity.
type ShipmentSchedule1Choice struct {

	// Specifies an earliest shipment date and a latest shipment date.
	ShipmentDateRange *ShipmentDateRange1 `xml:"ShipmntDtRg"`

	// Specifies a shipment schedule, ie, quantity that must be shipped no sooner than the earliest shipment date and no later than the latest shipment date.
	ShipmentSubSchedule []*ShipmentDateRange2 `xml:"ShipmntSubSchdl"`
}

func (s *ShipmentSchedule1Choice) AddShipmentDateRange() *ShipmentDateRange1 {
	s.ShipmentDateRange = new(ShipmentDateRange1)
	return s.ShipmentDateRange
}

func (s *ShipmentSchedule1Choice) AddShipmentSubSchedule() *ShipmentDateRange2 {
	newValue := new(ShipmentDateRange2)
	s.ShipmentSubSchedule = append(s.ShipmentSubSchedule, newValue)
	return newValue
}
