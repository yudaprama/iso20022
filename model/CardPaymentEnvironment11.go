package model

// Environment of the transaction given in a response to a request.
type CardPaymentEnvironment11 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification32 `xml:"AcqrrId,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType5 `xml:"PrtctdCardData,omitempty"`

	// Payment card performing the transaction.
	PlainCardData *PlainCardData3 `xml:"PlainCardData,omitempty"`
}

func (c *CardPaymentEnvironment11) AddAcquirerIdentification() *GenericIdentification32 {
	c.AcquirerIdentification = new(GenericIdentification32)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment11) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment11) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment11) AddProtectedCardData() *ContentInformationType5 {
	c.ProtectedCardData = new(ContentInformationType5)
	return c.ProtectedCardData
}

func (c *CardPaymentEnvironment11) AddPlainCardData() *PlainCardData3 {
	c.PlainCardData = new(PlainCardData3)
	return c.PlainCardData
}
