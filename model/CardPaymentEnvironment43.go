package model

// Environment of the diagnostic exchange.
type CardPaymentEnvironment43 struct {

	// Acquirer involved in the card payment transaction.
	Acquirer *Acquirer4 `xml:"Acqrr"`

	// Indicates if the acquirer is available to perform payment transactions.
	AcquirerAvailable *TrueFalseIndicator `xml:"AcqrrAvlbl,omitempty"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification53 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`
}

func (c *CardPaymentEnvironment43) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment43) SetAcquirerAvailable(value string) {
	c.AcquirerAvailable = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentEnvironment43) AddMerchantIdentification() *GenericIdentification53 {
	c.MerchantIdentification = new(GenericIdentification53)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment43) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}
