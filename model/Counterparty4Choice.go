package model

// Choice between buyer and seller.
type Counterparty4Choice struct {

	// Party that sells goods or services, or a financial instrument.
	Seller *PartyIdentificationAndAccount42 `xml:"Sellr"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *PartyIdentificationAndAccount42 `xml:"Buyr"`
}

func (c *Counterparty4Choice) AddSeller() *PartyIdentificationAndAccount42 {
	c.Seller = new(PartyIdentificationAndAccount42)
	return c.Seller
}

func (c *Counterparty4Choice) AddBuyer() *PartyIdentificationAndAccount42 {
	c.Buyer = new(PartyIdentificationAndAccount42)
	return c.Buyer
}
