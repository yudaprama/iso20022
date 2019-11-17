package model

// Specifies security rate details.
type CorporateActionRate17 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat11Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat11Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat12Choice `xml:"NewToOd,omitempty"`

	// Rate used to determine the cash consideration split across outturn settlement transactions that are the result of a transformation of the parent transaction.
	TransformationRate *PercentageRate `xml:"TrfrmatnRate,omitempty"`
}

func (c *CorporateActionRate17) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat11Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat11Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate17) AddAdditionalQuantityForExistingSecurities() *RatioFormat11Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat11Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate17) AddNewToOld() *RatioFormat12Choice {
	c.NewToOld = new(RatioFormat12Choice)
	return c.NewToOld
}

func (c *CorporateActionRate17) SetTransformationRate(value string) {
	c.TransformationRate = (*PercentageRate)(&value)
}
