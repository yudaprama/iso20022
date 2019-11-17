package model

// Choice between buyer and seller.
type Counterparty10Choice struct {

	// Party that sells goods or services, or a financial instrument.
	Seller *PartyIdentificationAndAccount131 `xml:"Sellr"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *PartyIdentificationAndAccount131 `xml:"Buyr"`
}

func (c *Counterparty10Choice) AddSeller() *PartyIdentificationAndAccount131 {
	c.Seller = new(PartyIdentificationAndAccount131)
	return c.Seller
}

func (c *Counterparty10Choice) AddBuyer() *PartyIdentificationAndAccount131 {
	c.Buyer = new(PartyIdentificationAndAccount131)
	return c.Buyer
}
