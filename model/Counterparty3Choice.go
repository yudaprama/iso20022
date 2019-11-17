package model

// Choice between buyer and seller.
type Counterparty3Choice struct {

	// Party that sells goods or services, or a financial instrument.
	Seller *PartyIdentificationAndAccount35 `xml:"Sellr"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *PartyIdentificationAndAccount35 `xml:"Buyr"`
}

func (c *Counterparty3Choice) AddSeller() *PartyIdentificationAndAccount35 {
	c.Seller = new(PartyIdentificationAndAccount35)
	return c.Seller
}

func (c *Counterparty3Choice) AddBuyer() *PartyIdentificationAndAccount35 {
	c.Buyer = new(PartyIdentificationAndAccount35)
	return c.Buyer
}
