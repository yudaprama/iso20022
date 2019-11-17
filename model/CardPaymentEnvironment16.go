package model

// Environment of the transaction given in a response to a request in a batch.
type CardPaymentEnvironment16 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification32 `xml:"Acqrr,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType5 `xml:"PrtctdCardData,omitempty"`

	// Payment card performing the transaction.
	PlainCardData *PlainCardData3 `xml:"PlainCardData,omitempty"`
}

func (c *CardPaymentEnvironment16) AddAcquirerIdentification() *GenericIdentification32 {
	c.AcquirerIdentification = new(GenericIdentification32)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment16) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment16) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment16) AddProtectedCardData() *ContentInformationType5 {
	c.ProtectedCardData = new(ContentInformationType5)
	return c.ProtectedCardData
}

func (c *CardPaymentEnvironment16) AddPlainCardData() *PlainCardData3 {
	c.PlainCardData = new(PlainCardData3)
	return c.PlainCardData
}
