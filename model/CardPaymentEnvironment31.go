package model

// Environment of the diagnostic exchange.
type CardPaymentEnvironment31 struct {

	// Acquirer involved in the card payment transaction.
	Acquirer *Acquirer2 `xml:"Acqrr"`

	// Indicates if the acquirer is available to perform payment transactions.
	AcquirerAvailable *TrueFalseIndicator `xml:"AcqrrAvlbl,omitempty"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`
}

func (c *CardPaymentEnvironment31) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment31) SetAcquirerAvailable(value string) {
	c.AcquirerAvailable = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentEnvironment31) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment31) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}
