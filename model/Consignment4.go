package model

// Specifies the arrangement of the transport of goods and services and the parties involved in this process.
type Consignment4 struct {

	// Party consigning goods as stipulated in the transport contract by the party ordering transport.
	Consignor *TradeParty3 `xml:"Consgnr,omitempty"`

	// Party to which goods are consigned.
	Consignee *TradeParty3 `xml:"Consgn,omitempty"`

	// Particular aircraft, vehicle, vessel or other device used for the transport of a consignment.
	TransportMeans []*TransportMeans3 `xml:"TrnsprtMeans,omitempty"`
}

func (c *Consignment4) AddConsignor() *TradeParty3 {
	c.Consignor = new(TradeParty3)
	return c.Consignor
}

func (c *Consignment4) AddConsignee() *TradeParty3 {
	c.Consignee = new(TradeParty3)
	return c.Consignee
}

func (c *Consignment4) AddTransportMeans() *TransportMeans3 {
	newValue := new(TransportMeans3)
	c.TransportMeans = append(c.TransportMeans, newValue)
	return newValue
}
