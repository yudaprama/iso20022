package model

// Specifies rate details.
type CorporateActionRate6 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat3Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat3Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat4Choice `xml:"NewToOd,omitempty"`

	// Quantity of new equities that will be derived by the exercise of a given quantity of intermediate securities.
	NewSecuritiesToUnderlyingSecurities *RatioFormat4Choice `xml:"NewSctiesToUndrlygScties,omitempty"`
}

func (c *CorporateActionRate6) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate6) AddAdditionalQuantityForExistingSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate6) AddNewToOld() *RatioFormat4Choice {
	c.NewToOld = new(RatioFormat4Choice)
	return c.NewToOld
}

func (c *CorporateActionRate6) AddNewSecuritiesToUnderlyingSecurities() *RatioFormat4Choice {
	c.NewSecuritiesToUnderlyingSecurities = new(RatioFormat4Choice)
	return c.NewSecuritiesToUnderlyingSecurities
}
