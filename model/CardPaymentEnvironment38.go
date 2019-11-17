package model

// Environment of the reconciliation for the acquirer.
type CardPaymentEnvironment38 struct {

	// Acquirer involved in the card payment reconciliation.
	AcquirerIdentification *GenericIdentification53 `xml:"AcqrrId"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI requesting the reconciliation.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`
}

func (c *CardPaymentEnvironment38) AddAcquirerIdentification() *GenericIdentification53 {
	c.AcquirerIdentification = new(GenericIdentification53)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment38) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment38) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}
