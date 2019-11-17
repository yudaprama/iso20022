package model

// Specifies security rate details.
type CorporateActionRate7 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat5Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat5Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat6Choice `xml:"NewToOd,omitempty"`

	// Quantity of new equities that will be derived by the exercise of a given quantity of intermediate securities.
	NewSecuritiesToUnderlyingSecurities *RatioFormat6Choice `xml:"NewSctiesToUndrlygScties,omitempty"`

	// Rate used to determine the cash consideration split across outturn settlement transactions that are the result of a transformation of the parent transaction.
	TransformationRate *PercentageRate `xml:"TrfrmatnRate,omitempty"`
}

func (c *CorporateActionRate7) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat5Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat5Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate7) AddAdditionalQuantityForExistingSecurities() *RatioFormat5Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat5Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate7) AddNewToOld() *RatioFormat6Choice {
	c.NewToOld = new(RatioFormat6Choice)
	return c.NewToOld
}

func (c *CorporateActionRate7) AddNewSecuritiesToUnderlyingSecurities() *RatioFormat6Choice {
	c.NewSecuritiesToUnderlyingSecurities = new(RatioFormat6Choice)
	return c.NewSecuritiesToUnderlyingSecurities
}

func (c *CorporateActionRate7) SetTransformationRate(value string) {
	c.TransformationRate = (*PercentageRate)(&value)
}
