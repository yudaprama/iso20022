package model

// Environment of the diagnostic exchange.
type CardPaymentEnvironment17 struct {

	// Version of acquirer configuration parameters.
	AcquirerParametersVersion *Max35Text `xml:"AcqrrParamsVrsn"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId"`
}

func (c *CardPaymentEnvironment17) SetAcquirerParametersVersion(value string) {
	c.AcquirerParametersVersion = (*Max35Text)(&value)
}

func (c *CardPaymentEnvironment17) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment17) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}
