package model

// Environment of the transaction.
type CardPaymentEnvironment25 struct {

	// Acquirer involved in the card payment reconciliation.
	Acquirer *Acquirer2 `xml:"Acqrr"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the reconciliation.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Data related to the components of the POI (Point Of Interaction) that have been performed the payment transactions.
	POIComponent []*PointOfInteractionComponent4 `xml:"POICmpnt,omitempty"`
}

func (c *CardPaymentEnvironment25) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment25) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment25) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment25) AddPOIComponent() *PointOfInteractionComponent4 {
	newValue := new(PointOfInteractionComponent4)
	c.POIComponent = append(c.POIComponent, newValue)
	return newValue
}
