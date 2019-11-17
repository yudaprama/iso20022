package model

// Provides details on the transportation of goods that are part of a commercial trade agreement.
type TransportDataSet3 struct {

	// Identifies the submitted transport data set.
	DataSetIdentification *DocumentIdentification1 `xml:"DataSetId"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *PartyIdentification26 `xml:"Buyr,omitempty"`

	// Party that sells goods or services, or a financial instrument.
	Seller *PartyIdentification26 `xml:"Sellr,omitempty"`

	// Party responsible for dispatching the goods.
	Consignor *PartyIdentification26 `xml:"Consgnr"`

	// Party to whom the goods must be delivered.
	Consignee *PartyIdentification26 `xml:"Consgn,omitempty"`

	// Party to whom the goods must be delivered in the end.
	ShipTo *PartyIdentification26 `xml:"ShipTo,omitempty"`

	// Specifies the shipment date, the charges, the routing and the goods that are described in the transport document.
	TransportInformation *TransportDetails2 `xml:"TrnsprtInf"`
}

func (t *TransportDataSet3) AddDataSetIdentification() *DocumentIdentification1 {
	t.DataSetIdentification = new(DocumentIdentification1)
	return t.DataSetIdentification
}

func (t *TransportDataSet3) AddBuyer() *PartyIdentification26 {
	t.Buyer = new(PartyIdentification26)
	return t.Buyer
}

func (t *TransportDataSet3) AddSeller() *PartyIdentification26 {
	t.Seller = new(PartyIdentification26)
	return t.Seller
}

func (t *TransportDataSet3) AddConsignor() *PartyIdentification26 {
	t.Consignor = new(PartyIdentification26)
	return t.Consignor
}

func (t *TransportDataSet3) AddConsignee() *PartyIdentification26 {
	t.Consignee = new(PartyIdentification26)
	return t.Consignee
}

func (t *TransportDataSet3) AddShipTo() *PartyIdentification26 {
	t.ShipTo = new(PartyIdentification26)
	return t.ShipTo
}

func (t *TransportDataSet3) AddTransportInformation() *TransportDetails2 {
	t.TransportInformation = new(TransportDetails2)
	return t.TransportInformation
}
