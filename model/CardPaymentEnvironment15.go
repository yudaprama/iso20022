package model

// Environment of the reconciliation exchange.
type CardPaymentEnvironment15 struct {

	// Acquirer involved in the card payment reconciliation.
	Acquirer *Acquirer2 `xml:"Acqrr"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI requesting the reconciliation.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`
}

func (c *CardPaymentEnvironment15) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment15) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment15) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}
