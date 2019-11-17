package model

// Information on the shipment date, the charges, the routing and the goods described in the transport document.
type TransportDetails3 struct {

	// Reference to the identification of  the underlying transport document.
	TransportDocumentReference []*DocumentIdentification7 `xml:"TrnsprtDocRef"`

	// Goods that are transported.
	TransportedGoods []*TransportedGoods1 `xml:"TrnsprtdGoods"`

	// Physical packaging of goods for transport.
	Consignment *Consignment3 `xml:"Consgnmt,omitempty"`

	// Information related to the conveyance of goods.
	RoutingSummary *TransportMeans4 `xml:"RtgSummry"`

	// Shipment date, actual or proposed.
	ShipmentDate *ShipmentDate1Choice `xml:"ShipmntDt"`

	// Charges related to the conveyance of goods.
	FreightCharges *Charge25 `xml:"FrghtChrgs,omitempty"`

	// Specifies the applicable Incoterm and associated location.
	Incoterms *Incoterms4 `xml:"Incotrms,omitempty"`
}

func (t *TransportDetails3) AddTransportDocumentReference() *DocumentIdentification7 {
	newValue := new(DocumentIdentification7)
	t.TransportDocumentReference = append(t.TransportDocumentReference, newValue)
	return newValue
}

func (t *TransportDetails3) AddTransportedGoods() *TransportedGoods1 {
	newValue := new(TransportedGoods1)
	t.TransportedGoods = append(t.TransportedGoods, newValue)
	return newValue
}

func (t *TransportDetails3) AddConsignment() *Consignment3 {
	t.Consignment = new(Consignment3)
	return t.Consignment
}

func (t *TransportDetails3) AddRoutingSummary() *TransportMeans4 {
	t.RoutingSummary = new(TransportMeans4)
	return t.RoutingSummary
}

func (t *TransportDetails3) AddShipmentDate() *ShipmentDate1Choice {
	t.ShipmentDate = new(ShipmentDate1Choice)
	return t.ShipmentDate
}

func (t *TransportDetails3) AddFreightCharges() *Charge25 {
	t.FreightCharges = new(Charge25)
	return t.FreightCharges
}

func (t *TransportDetails3) AddIncoterms() *Incoterms4 {
	t.Incoterms = new(Incoterms4)
	return t.Incoterms
}
