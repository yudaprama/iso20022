package model

// Specifies periods of a corporate action.
type CorporateActionPeriod10 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period3Choice `xml:"PricClctnPrd,omitempty"`

	// Period during which the interest rate has been applied.
	InterestPeriod *Period3Choice `xml:"IntrstPrd,omitempty"`

	// Period during a take-over where any outstanding equity must be purchased by the take-over company.
	CompulsoryPurchasePeriod *Period3Choice `xml:"CmplsryPurchsPrd,omitempty"`

	// Period during which the security is blocked.
	BlockingPeriod *Period3Choice `xml:"BlckgPrd,omitempty"`

	// Period assigned by the court in a class action. It determines the client's eligible transactions that will be included in the class action and used to determine the resulting entitlement.
	ClaimPeriod *Period3Choice `xml:"ClmPrd,omitempty"`

	// Period defining the last date for which book entry transfers will be accepted and the date on which the suspension will be released and book entry transfer processing will resume.
	DepositorySuspensionPeriodForBookEntryTransfer *Period3Choice `xml:"DpstrySspnsnPrdForBookNtryTrf,omitempty"`

	// Period defining the last date for which deposits, into nominee name, at the agent will be accepted and the date on which the suspension will be released and deposits at agent will resume.
	DepositorySuspensionPeriodForDepositAtAgent *Period3Choice `xml:"DpstrySspnsnPrdForDpstAtAgt,omitempty"`

	// Period defining the last date for which deposits will be accepted and the date on which the suspension will be released and deposits will resume.
	DepositorySuspensionPeriodForDeposit *Period3Choice `xml:"DpstrySspnsnPrdForDpst,omitempty"`

	// Period defining the last date for which pledges will be accepted and the date on which the suspension will be released and pledge processing will resume.
	DepositorySuspensionPeriodForPledge *Period3Choice `xml:"DpstrySspnsnPrdForPldg,omitempty"`

	// Period defining the last date for which intra-position balances can be segregated and the date on which the suspension will be released and the ability to segregate intra-position balances will resume.
	DepositorySuspensionPeriodForSegregation *Period3Choice `xml:"DpstrySspnsnPrdForSgrtn,omitempty"`

	// Period defining the last date for which withdrawals, from nominee name at the agent will be accepted and the date on which the suspension will be released and withdrawals at agent processing will resume.
	DepositorySuspensionPeriodForWithdrawalAtAgent *Period3Choice `xml:"DpstrySspnsnPrdForWdrwlAtAgt,omitempty"`

	// Period defining the last date for which physical withdrawals in the nominee's name will be accepted and the date on which the suspension will be released and physical withdrawals in the nominee's name will resume.
	DepositorySuspensionPeriodForWithdrawalInNomineeName *Period3Choice `xml:"DpstrySspnsnPrdForWdrwlInNmneeNm,omitempty"`

	// Period defining the last date on which withdrawal requests in street name's will be accepted on the event security and the date on which the suspension will be released and withdrawal in street name's processing on the event security will resume.
	DepositorySuspensionPeriodForWithdrawalInStreetName *Period3Choice `xml:"DpstrySspnsnPrdForWdrwlInStrtNm,omitempty"`

	// Period defining the last date on which shareholder registration will be accepted by the issuer and the date on which shareholder registration will resume.
	BookClosurePeriod *Period3Choice `xml:"BookClsrPrd,omitempty"`

	// Period during which the settlement activities at the co-depositories are suspended in order to stabilise the holdings at the CSD.
	CoDepositoriesSuspensionPeriod *Period3Choice `xml:"CoDpstriesSspnsnPrd,omitempty"`

	// Period during which a physical certificate can be split.
	SplitPeriod *Period3Choice `xml:"SpltPrd,omitempty"`
}

func (c *CorporateActionPeriod10) AddPriceCalculationPeriod() *Period3Choice {
	c.PriceCalculationPeriod = new(Period3Choice)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod10) AddInterestPeriod() *Period3Choice {
	c.InterestPeriod = new(Period3Choice)
	return c.InterestPeriod
}

func (c *CorporateActionPeriod10) AddCompulsoryPurchasePeriod() *Period3Choice {
	c.CompulsoryPurchasePeriod = new(Period3Choice)
	return c.CompulsoryPurchasePeriod
}

func (c *CorporateActionPeriod10) AddBlockingPeriod() *Period3Choice {
	c.BlockingPeriod = new(Period3Choice)
	return c.BlockingPeriod
}

func (c *CorporateActionPeriod10) AddClaimPeriod() *Period3Choice {
	c.ClaimPeriod = new(Period3Choice)
	return c.ClaimPeriod
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForBookEntryTransfer() *Period3Choice {
	c.DepositorySuspensionPeriodForBookEntryTransfer = new(Period3Choice)
	return c.DepositorySuspensionPeriodForBookEntryTransfer
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForDepositAtAgent() *Period3Choice {
	c.DepositorySuspensionPeriodForDepositAtAgent = new(Period3Choice)
	return c.DepositorySuspensionPeriodForDepositAtAgent
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForDeposit() *Period3Choice {
	c.DepositorySuspensionPeriodForDeposit = new(Period3Choice)
	return c.DepositorySuspensionPeriodForDeposit
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForPledge() *Period3Choice {
	c.DepositorySuspensionPeriodForPledge = new(Period3Choice)
	return c.DepositorySuspensionPeriodForPledge
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForSegregation() *Period3Choice {
	c.DepositorySuspensionPeriodForSegregation = new(Period3Choice)
	return c.DepositorySuspensionPeriodForSegregation
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForWithdrawalAtAgent() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawalAtAgent = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawalAtAgent
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForWithdrawalInNomineeName() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawalInNomineeName = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawalInNomineeName
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForWithdrawalInStreetName() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawalInStreetName = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawalInStreetName
}

func (c *CorporateActionPeriod10) AddBookClosurePeriod() *Period3Choice {
	c.BookClosurePeriod = new(Period3Choice)
	return c.BookClosurePeriod
}

func (c *CorporateActionPeriod10) AddCoDepositoriesSuspensionPeriod() *Period3Choice {
	c.CoDepositoriesSuspensionPeriod = new(Period3Choice)
	return c.CoDepositoriesSuspensionPeriod
}

func (c *CorporateActionPeriod10) AddSplitPeriod() *Period3Choice {
	c.SplitPeriod = new(Period3Choice)
	return c.SplitPeriod
}
