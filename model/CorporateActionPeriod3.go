package model

// Specifies periods of a corporate action.
type CorporateActionPeriod3 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period1Choice `xml:"PricClctnPrd,omitempty"`

	// Period during which the interest rate has been applied.
	InterestPeriod *Period1Choice `xml:"IntrstPrd,omitempty"`

	// Period during a take-over where any outstanding equity must be purchased by the take-over company.
	CompulsoryPurchasePeriod *Period1Choice `xml:"CmplsryPurchsPrd,omitempty"`

	// Period during which the security is blocked.
	BlockingPeriod *Period1Choice `xml:"BlckgPrd,omitempty"`

	// Period assigned by the court in a class action. It determines the client's eligible transactions that will be included in the class action and used to determine the resulting entitlement.
	ClaimPeriod *Period1Choice `xml:"ClmPrd,omitempty"`

	// Period defining the last date for which book entry transfers will be accepted and the date on which the suspension will be released and book entry transfer processing will resume.
	DepositorySuspensionPeriodForBookEntryTransfer *Period1Choice `xml:"DpstrySspnsnPrdForBookNtryTrf,omitempty"`

	// Period defining the last date for which deposits, into nominee name, at the agent will be accepted and the date on which the suspension will be released and deposits at agent will resume.
	DepositorySuspensionPeriodForDepositAtAgent *Period1Choice `xml:"DpstrySspnsnPrdForDpstAtAgt,omitempty"`

	// Period defining the last date for which deposits will be accepted and the date on which the suspension will be released and deposits will resume.
	DepositorySuspensionPeriodForDeposit *Period1Choice `xml:"DpstrySspnsnPrdForDpst,omitempty"`

	// Period defining the last date for which pledges will be accepted and the date on which the suspension will be released and pledge processing will resume.
	DepositorySuspensionPeriodForPledge *Period1Choice `xml:"DpstrySspnsnPrdForPldg,omitempty"`

	// Period defining the last date for which intra-position balances can be segregated and the date on which the suspension will be released and the ability to segregate intra-position balances will resume.
	DepositorySuspensionPeriodForSegregation *Period1Choice `xml:"DpstrySspnsnPrdForSgrtn,omitempty"`

	// Period defining the last date for which withdrawals, from nominee name at the agent will be accepted and the date on which the suspension will be released and withdrawals at agent processing will resume.
	DepositorySuspensionPeriodForWithdrawalAtAgent *Period1Choice `xml:"DpstrySspnsnPrdForWdrwlAtAgt,omitempty"`

	// Period defining the last date for which physical withdrawals in the nominee's name will be accepted and the date on which the suspension will be released and physical withdrawals in the nominee's name will resume.
	DepositorySuspensionPeriodForWithdrawalInNomineeName *Period1Choice `xml:"DpstrySspnsnPrdForWdrwlInNmneeNm,omitempty"`

	// Period defining the last date on which withdrawal requests in street name's will be accepted on the event security and the date on which the suspension will be released and withdrawal in street name's processing on the event security will resume.
	DepositorySuspensionPeriodForWithdrawalInStreetName *Period1Choice `xml:"DpstrySspnsnPrdForWdrwlInStrtNm,omitempty"`

	// Period defining the last date on which shareholder registration will be accepted by the issuer and the date on which shareholder registration will resume.
	BookClosurePeriod *Period1Choice `xml:"BookClsrPrd,omitempty"`
}

func (c *CorporateActionPeriod3) AddPriceCalculationPeriod() *Period1Choice {
	c.PriceCalculationPeriod = new(Period1Choice)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod3) AddInterestPeriod() *Period1Choice {
	c.InterestPeriod = new(Period1Choice)
	return c.InterestPeriod
}

func (c *CorporateActionPeriod3) AddCompulsoryPurchasePeriod() *Period1Choice {
	c.CompulsoryPurchasePeriod = new(Period1Choice)
	return c.CompulsoryPurchasePeriod
}

func (c *CorporateActionPeriod3) AddBlockingPeriod() *Period1Choice {
	c.BlockingPeriod = new(Period1Choice)
	return c.BlockingPeriod
}

func (c *CorporateActionPeriod3) AddClaimPeriod() *Period1Choice {
	c.ClaimPeriod = new(Period1Choice)
	return c.ClaimPeriod
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForBookEntryTransfer() *Period1Choice {
	c.DepositorySuspensionPeriodForBookEntryTransfer = new(Period1Choice)
	return c.DepositorySuspensionPeriodForBookEntryTransfer
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForDepositAtAgent() *Period1Choice {
	c.DepositorySuspensionPeriodForDepositAtAgent = new(Period1Choice)
	return c.DepositorySuspensionPeriodForDepositAtAgent
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForDeposit() *Period1Choice {
	c.DepositorySuspensionPeriodForDeposit = new(Period1Choice)
	return c.DepositorySuspensionPeriodForDeposit
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForPledge() *Period1Choice {
	c.DepositorySuspensionPeriodForPledge = new(Period1Choice)
	return c.DepositorySuspensionPeriodForPledge
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForSegregation() *Period1Choice {
	c.DepositorySuspensionPeriodForSegregation = new(Period1Choice)
	return c.DepositorySuspensionPeriodForSegregation
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForWithdrawalAtAgent() *Period1Choice {
	c.DepositorySuspensionPeriodForWithdrawalAtAgent = new(Period1Choice)
	return c.DepositorySuspensionPeriodForWithdrawalAtAgent
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForWithdrawalInNomineeName() *Period1Choice {
	c.DepositorySuspensionPeriodForWithdrawalInNomineeName = new(Period1Choice)
	return c.DepositorySuspensionPeriodForWithdrawalInNomineeName
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForWithdrawalInStreetName() *Period1Choice {
	c.DepositorySuspensionPeriodForWithdrawalInStreetName = new(Period1Choice)
	return c.DepositorySuspensionPeriodForWithdrawalInStreetName
}

func (c *CorporateActionPeriod3) AddBookClosurePeriod() *Period1Choice {
	c.BookClosurePeriod = new(Period1Choice)
	return c.BookClosurePeriod
}
