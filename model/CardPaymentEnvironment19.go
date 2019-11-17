package model

// Environment of the reconciliation for the acquirer.
type CardPaymentEnvironment19 struct {

	// Acquirer involved in the card payment reconciliation.
	AcquirerIdentification *GenericIdentification32 `xml:"AcqrrId"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI requesting the reconciliation.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`
}

func (c *CardPaymentEnvironment19) AddAcquirerIdentification() *GenericIdentification32 {
	c.AcquirerIdentification = new(GenericIdentification32)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment19) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment19) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}
