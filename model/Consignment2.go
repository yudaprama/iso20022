package model

// Specifies the arrangement of the transport of goods and services and the parties involved in this process.
type Consignment2 struct {

	// Party consigning goods as stipulated in the transport contract by the party ordering transport.
	Consignor *TradeParty1 `xml:"Consgnr,omitempty"`

	// Party to which goods are consigned.
	Consignee *TradeParty1 `xml:"Consgn,omitempty"`

	// Particular aircraft, vehicle, vessel or other device used for the transport of a consignment.
	TransportMeans []*TransportMeans3 `xml:"TrnsprtMeans,omitempty"`
}

func (c *Consignment2) AddConsignor() *TradeParty1 {
	c.Consignor = new(TradeParty1)
	return c.Consignor
}

func (c *Consignment2) AddConsignee() *TradeParty1 {
	c.Consignee = new(TradeParty1)
	return c.Consignee
}

func (c *Consignment2) AddTransportMeans() *TransportMeans3 {
	newValue := new(TransportMeans3)
	c.TransportMeans = append(c.TransportMeans, newValue)
	return newValue
}
