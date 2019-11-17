package model

// Environment of the transaction given in a response to a request.
type CardPaymentEnvironment21 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification32 `xml:"AcqrrId,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType7 `xml:"PrtctdCardData,omitempty"`

	// Payment card performing the transaction.
	PlainCardData *PlainCardData5 `xml:"PlainCardData,omitempty"`
}

func (c *CardPaymentEnvironment21) AddAcquirerIdentification() *GenericIdentification32 {
	c.AcquirerIdentification = new(GenericIdentification32)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment21) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment21) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment21) AddProtectedCardData() *ContentInformationType7 {
	c.ProtectedCardData = new(ContentInformationType7)
	return c.ProtectedCardData
}

func (c *CardPaymentEnvironment21) AddPlainCardData() *PlainCardData5 {
	c.PlainCardData = new(PlainCardData5)
	return c.PlainCardData
}
