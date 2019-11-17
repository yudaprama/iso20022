package model

// Environment of the transaction given in a response to a request.
type CardPaymentEnvironment3 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer1 `xml:"Acqrr,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType1 `xml:"PrtctdCardData,omitempty"`

	// Payment card performing the transaction.
	PlainCardData *PlainCardData3 `xml:"PlainCardData,omitempty"`
}

func (c *CardPaymentEnvironment3) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment3) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment3) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment3) AddProtectedCardData() *ContentInformationType1 {
	c.ProtectedCardData = new(ContentInformationType1)
	return c.ProtectedCardData
}

func (c *CardPaymentEnvironment3) AddPlainCardData() *PlainCardData3 {
	c.PlainCardData = new(PlainCardData3)
	return c.PlainCardData
}
