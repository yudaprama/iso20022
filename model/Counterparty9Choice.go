package model

// Choice between buyer and seller.
type Counterparty9Choice struct {

	// Party that sells goods or services, or a financial instrument.
	Seller *PartyIdentificationAndAccount106 `xml:"Sellr"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *PartyIdentificationAndAccount106 `xml:"Buyr"`
}

func (c *Counterparty9Choice) AddSeller() *PartyIdentificationAndAccount106 {
	c.Seller = new(PartyIdentificationAndAccount106)
	return c.Seller
}

func (c *Counterparty9Choice) AddBuyer() *PartyIdentificationAndAccount106 {
	c.Buyer = new(PartyIdentificationAndAccount106)
	return c.Buyer
}
