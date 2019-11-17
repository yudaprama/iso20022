package model

// Environment of the reconciliation exchange.
type CardPaymentEnvironment7 struct {

	// Acquirer involved in the card payment reconciliation.
	Acquirer *Acquirer1 `xml:"Acqrr"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI requesting the reconciliation.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`
}

func (c *CardPaymentEnvironment7) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment7) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment7) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}
