package model

// Environment of the transaction.
type CardPaymentEnvironment37 struct {

	// Acquirer involved in the card payment reconciliation.
	Acquirer *Acquirer4 `xml:"Acqrr"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification53 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the reconciliation.
	POIIdentification *GenericIdentification53 `xml:"POIId,omitempty"`

	// Data related to the components of the POI (Point Of Interaction) that have been performed the payment transactions.
	POIComponent []*PointOfInteractionComponent5 `xml:"POICmpnt,omitempty"`
}

func (c *CardPaymentEnvironment37) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment37) AddMerchantIdentification() *GenericIdentification53 {
	c.MerchantIdentification = new(GenericIdentification53)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment37) AddPOIIdentification() *GenericIdentification53 {
	c.POIIdentification = new(GenericIdentification53)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment37) AddPOIComponent() *PointOfInteractionComponent5 {
	newValue := new(PointOfInteractionComponent5)
	c.POIComponent = append(c.POIComponent, newValue)
	return newValue
}
