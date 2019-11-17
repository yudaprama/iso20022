package model

// Specifies periods related to a corporate action option.
type CorporateActionPeriod5 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period1Choice `xml:"PricClctnPrd,omitempty"`

	// Period during which both old and new equity may be traded simultaneously, for example, consolidation of equity or splitting of equity.
	ParallelTradingPeriod *Period1Choice `xml:"ParllTradgPrd,omitempty"`

	// Period during which the specified option, or all options of the event, remains valid, for example, offer period.
	ActionPeriod *Period1Choice `xml:"ActnPrd,omitempty"`

	// Period during which the shareholder can revoke, change or withdraw its instruction.
	RevocabilityPeriod *Period1Choice `xml:"RvcbltyPrd,omitempty"`

	// Period during which the privilege is not available, for example, this can happen whenever a meeting takes place or whenever a coupon payment is due.
	PrivilegeSuspensionPeriod *Period1Choice `xml:"PrvlgSspnsnPrd,omitempty"`

	// Period during which the participant of the account servicer can revoke change or withdraw its instructions.
	AccountServicerRevocabilityPeriod *Period1Choice `xml:"AcctSvcrRvcbltyPrd,omitempty"`

	// Period defining the last date on which withdrawal in street name requests on the outturn security will be accepted and the date on which the suspension will be released and withdrawal by transfer processing on the outturn security will resume.
	DepositorySuspensionPeriodForWithdrawal *Period1Choice `xml:"DpstrySspnsnPrdForWdrwl,omitempty"`
}

func (c *CorporateActionPeriod5) AddPriceCalculationPeriod() *Period1Choice {
	c.PriceCalculationPeriod = new(Period1Choice)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod5) AddParallelTradingPeriod() *Period1Choice {
	c.ParallelTradingPeriod = new(Period1Choice)
	return c.ParallelTradingPeriod
}

func (c *CorporateActionPeriod5) AddActionPeriod() *Period1Choice {
	c.ActionPeriod = new(Period1Choice)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod5) AddRevocabilityPeriod() *Period1Choice {
	c.RevocabilityPeriod = new(Period1Choice)
	return c.RevocabilityPeriod
}

func (c *CorporateActionPeriod5) AddPrivilegeSuspensionPeriod() *Period1Choice {
	c.PrivilegeSuspensionPeriod = new(Period1Choice)
	return c.PrivilegeSuspensionPeriod
}

func (c *CorporateActionPeriod5) AddAccountServicerRevocabilityPeriod() *Period1Choice {
	c.AccountServicerRevocabilityPeriod = new(Period1Choice)
	return c.AccountServicerRevocabilityPeriod
}

func (c *CorporateActionPeriod5) AddDepositorySuspensionPeriodForWithdrawal() *Period1Choice {
	c.DepositorySuspensionPeriodForWithdrawal = new(Period1Choice)
	return c.DepositorySuspensionPeriodForWithdrawal
}
