package model

// Environment of the diagnostic exchange.
type CardPaymentEnvironment8 struct {

	// Version of the payment application configuration parameters of the POI.
	AcquirerParametersVersion *ISODateTime `xml:"AcqrrParamsVrsn"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId"`
}

func (c *CardPaymentEnvironment8) SetAcquirerParametersVersion(value string) {
	c.AcquirerParametersVersion = (*ISODateTime)(&value)
}

func (c *CardPaymentEnvironment8) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment8) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}
