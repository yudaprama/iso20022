package model

// Specifies how the supply chain shipping arrangements and the delivery of products and/or services as well as related documentation.
type TradeDelivery2 struct {

	// Actual delivery period of the products and/or services.
	DeliveryPeriod *Period1 `xml:"DlvryPrd,omitempty"`

	// Actual delivery date/time of the products and/or services.
	DeliveryDateTime *ISODateTime `xml:"DlvryDtTm,omitempty"`

	// Party from whom the goods are dispatched.
	ShipFrom *TradeParty3 `xml:"ShipFr,omitempty"`

	// Party to whom the goods are dispatched.
	ShipTo *TradeParty3 `xml:"ShipTo,omitempty"`

	// Final party to whom the goods are dispatched.
	UltimateShipTo *TradeParty3 `xml:"UltmtShipTo,omitempty"`

	// Delivery note related to the delivery of the products and/or services.
	DeliveryNote *DocumentIdentification22 `xml:"DlvryNote,omitempty"`

	// Physical consolidation of goods for transport.
	Consignment []*Consignment4 `xml:"Consgnmt,omitempty"`
}

func (t *TradeDelivery2) AddDeliveryPeriod() *Period1 {
	t.DeliveryPeriod = new(Period1)
	return t.DeliveryPeriod
}

func (t *TradeDelivery2) SetDeliveryDateTime(value string) {
	t.DeliveryDateTime = (*ISODateTime)(&value)
}

func (t *TradeDelivery2) AddShipFrom() *TradeParty3 {
	t.ShipFrom = new(TradeParty3)
	return t.ShipFrom
}

func (t *TradeDelivery2) AddShipTo() *TradeParty3 {
	t.ShipTo = new(TradeParty3)
	return t.ShipTo
}

func (t *TradeDelivery2) AddUltimateShipTo() *TradeParty3 {
	t.UltimateShipTo = new(TradeParty3)
	return t.UltimateShipTo
}

func (t *TradeDelivery2) AddDeliveryNote() *DocumentIdentification22 {
	t.DeliveryNote = new(DocumentIdentification22)
	return t.DeliveryNote
}

func (t *TradeDelivery2) AddConsignment() *Consignment4 {
	newValue := new(Consignment4)
	t.Consignment = append(t.Consignment, newValue)
	return newValue
}
