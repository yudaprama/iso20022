package model

// Choice between proposed and actual shipment date.
type ShipmentDate1Choice struct {

	// Proposed date on which the goods should be shipped.
	ProposedShipmentDate *ISODate `xml:"PropsdShipmntDt"`

	// Actual date whereby the goods were shipped.
	ActualShipmentDate *ISODate `xml:"ActlShipmntDt"`
}

func (s *ShipmentDate1Choice) SetProposedShipmentDate(value string) {
	s.ProposedShipmentDate = (*ISODate)(&value)
}

func (s *ShipmentDate1Choice) SetActualShipmentDate(value string) {
	s.ActualShipmentDate = (*ISODate)(&value)
}
