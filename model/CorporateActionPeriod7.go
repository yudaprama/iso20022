package model

// Specifies periods related to a corporate action option.
type CorporateActionPeriod7 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period3Choice `xml:"PricClctnPrd,omitempty"`

	// Period during which both old and new equity may be traded simultaneously, for example, consolidation of equity or splitting of equity.
	ParallelTradingPeriod *Period3Choice `xml:"ParllTradgPrd,omitempty"`

	// Period during which the specified option, or all options of the event, remains valid, for example, offer period.
	ActionPeriod *Period3Choice `xml:"ActnPrd,omitempty"`

	// Period during which the shareholder can revoke, change or withdraw its instruction.
	RevocabilityPeriod *Period3Choice `xml:"RvcbltyPrd,omitempty"`

	// Period during which the privilege is not available, for example, this can happen whenever a meeting takes place or whenever a coupon payment is due.
	PrivilegeSuspensionPeriod *Period3Choice `xml:"PrvlgSspnsnPrd,omitempty"`

	// Period during which the participant of the account servicer can revoke change or withdraw its instructions.
	AccountServicerRevocabilityPeriod *Period3Choice `xml:"AcctSvcrRvcbltyPrd,omitempty"`

	// Period defining the last date on which withdrawal in street name requests on the outturn security will be accepted and the date on which the suspension will be released and withdrawal by transfer processing on the outturn security will resume.
	DepositorySuspensionPeriodForWithdrawal *Period3Choice `xml:"DpstrySspnsnPrdForWdrwl,omitempty"`
}

func (c *CorporateActionPeriod7) AddPriceCalculationPeriod() *Period3Choice {
	c.PriceCalculationPeriod = new(Period3Choice)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod7) AddParallelTradingPeriod() *Period3Choice {
	c.ParallelTradingPeriod = new(Period3Choice)
	return c.ParallelTradingPeriod
}

func (c *CorporateActionPeriod7) AddActionPeriod() *Period3Choice {
	c.ActionPeriod = new(Period3Choice)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod7) AddRevocabilityPeriod() *Period3Choice {
	c.RevocabilityPeriod = new(Period3Choice)
	return c.RevocabilityPeriod
}

func (c *CorporateActionPeriod7) AddPrivilegeSuspensionPeriod() *Period3Choice {
	c.PrivilegeSuspensionPeriod = new(Period3Choice)
	return c.PrivilegeSuspensionPeriod
}

func (c *CorporateActionPeriod7) AddAccountServicerRevocabilityPeriod() *Period3Choice {
	c.AccountServicerRevocabilityPeriod = new(Period3Choice)
	return c.AccountServicerRevocabilityPeriod
}

func (c *CorporateActionPeriod7) AddDepositorySuspensionPeriodForWithdrawal() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawal = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawal
}
