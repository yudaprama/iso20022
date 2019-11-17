package model

// Environment of the diagnostic exchange.
type CardPaymentEnvironment29 struct {

	// Acquirer involved in the card payment transaction.
	Acquirer *Acquirer2 `xml:"Acqrr"`

	// The availability of the acquirer to process transaction must be provided.
	AcquirerAvailabilityRequested *TrueFalseIndicator `xml:"AcqrrAvlbtyReqd,omitempty"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Data related to the components of the POI (Point Of Interaction) performing the payment transactions.
	POIComponent []*PointOfInteractionComponent4 `xml:"POICmpnt,omitempty"`
}

func (c *CardPaymentEnvironment29) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment29) SetAcquirerAvailabilityRequested(value string) {
	c.AcquirerAvailabilityRequested = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentEnvironment29) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment29) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment29) AddPOIComponent() *PointOfInteractionComponent4 {
	newValue := new(PointOfInteractionComponent4)
	c.POIComponent = append(c.POIComponent, newValue)
	return newValue
}
