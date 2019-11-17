package model

// Choice between buyer and seller.
type Counterparty1Choice struct {

	// Party that sells goods or services, or a financial instrument.
	Seller *PartyIdentificationAndAccount1 `xml:"Sellr"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *PartyIdentificationAndAccount1 `xml:"Buyr"`
}

func (c *Counterparty1Choice) AddSeller() *PartyIdentificationAndAccount1 {
	c.Seller = new(PartyIdentificationAndAccount1)
	return c.Seller
}

func (c *Counterparty1Choice) AddBuyer() *PartyIdentificationAndAccount1 {
	c.Buyer = new(PartyIdentificationAndAccount1)
	return c.Buyer
}
