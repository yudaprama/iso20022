package model

// Specifies rate details.
type CorporateActionRate21 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat3Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat3Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat4Choice `xml:"NewToOd,omitempty"`
}

func (c *CorporateActionRate21) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate21) AddAdditionalQuantityForExistingSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate21) AddNewToOld() *RatioFormat4Choice {
	c.NewToOld = new(RatioFormat4Choice)
	return c.NewToOld
}
