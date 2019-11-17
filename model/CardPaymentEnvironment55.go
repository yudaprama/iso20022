package model

// Environment of the diagnostic exchange.
type CardPaymentEnvironment55 struct {

	// Acquirer involved in the card payment transaction.
	Acquirer *Acquirer4 `xml:"Acqrr"`

	// The availability of the acquirer to process transaction must be provided.
	AcquirerAvailabilityRequested *TrueFalseIndicator `xml:"AcqrrAvlbtyReqd,omitempty"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification53 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Data related to the components of the POI (Point Of Interaction) performing the payment transactions.
	POIComponent []*PointOfInteractionComponent6 `xml:"POICmpnt,omitempty"`
}

func (c *CardPaymentEnvironment55) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment55) SetAcquirerAvailabilityRequested(value string) {
	c.AcquirerAvailabilityRequested = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentEnvironment55) AddMerchantIdentification() *GenericIdentification53 {
	c.MerchantIdentification = new(GenericIdentification53)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment55) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment55) AddPOIComponent() *PointOfInteractionComponent6 {
	newValue := new(PointOfInteractionComponent6)
	c.POIComponent = append(c.POIComponent, newValue)
	return newValue
}
