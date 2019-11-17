package model

// Further details on the shipment conditions.
type ShipmentAttribute1 struct {

	// Shipment conditions.
	Conditions *ExternalShipmentCondition1Code `xml:"Conds,omitempty"`

	// Expected shipment date.
	ExpectedDate *ISODate `xml:"XpctdDt,omitempty"`

	// Country in which the counter party is located.
	CountryOfCounterParty *CountryCode `xml:"CtryOfCntrPty"`
}

func (s *ShipmentAttribute1) SetConditions(value string) {
	s.Conditions = (*ExternalShipmentCondition1Code)(&value)
}

func (s *ShipmentAttribute1) SetExpectedDate(value string) {
	s.ExpectedDate = (*ISODate)(&value)
}

func (s *ShipmentAttribute1) SetCountryOfCounterParty(value string) {
	s.CountryOfCounterParty = (*CountryCode)(&value)
}
