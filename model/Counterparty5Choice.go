package model

// Choice between buyer and seller.
type Counterparty5Choice struct {

	// Party that sells goods or services, or a financial instrument.
	Seller *PartyIdentificationAndAccount45 `xml:"Sellr"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *PartyIdentificationAndAccount45 `xml:"Buyr"`
}

func (c *Counterparty5Choice) AddSeller() *PartyIdentificationAndAccount45 {
	c.Seller = new(PartyIdentificationAndAccount45)
	return c.Seller
}

func (c *Counterparty5Choice) AddBuyer() *PartyIdentificationAndAccount45 {
	c.Buyer = new(PartyIdentificationAndAccount45)
	return c.Buyer
}
