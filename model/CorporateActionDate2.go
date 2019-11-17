package model

// Specifies corporate action dates.
type CorporateActionDate2 struct {

	// Date on which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat4Choice `xml:"RcrdDt,omitempty"`

	// Date on which a process is to be completed or becomes effective.
	EffectiveDate *DateFormat4Choice `xml:"FctvDt,omitempty"`

	// Last day a holder can deliver the securities that it had previously protected.
	CoverExpirationDate *DateFormat4Choice `xml:"CoverXprtnDt,omitempty"`

	// Date on which all or part of any holding bought in a unit trust is subject to being treated as capital rather than income. This is normally one day after the previous distribution's ex date.
	EqualisationDate *DateFormat4Choice `xml:"EqulstnDt,omitempty"`

	// Date/time at which the margin rate will be determined .
	MarginFixingDate *DateFormat4Choice `xml:"MrgnFxgDt,omitempty"`

	// Date on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat4Choice `xml:"LtryDt,omitempty"`

	// Last date a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat4Choice `xml:"PrtctDt,omitempty"`

	// Date upon which the terms of the take-over become unconditional as to acceptances.
	UnconditionalDate *DateFormat4Choice `xml:"UcondlDt,omitempty"`

	// Date on which all conditions, including regulatory, legal etc. pertaining to the take-over, have been met.
	WhollyUnconditionalDate *DateFormat4Choice `xml:"WhlyUcondlDt,omitempty"`

	// Date on which results are published, eg, results of an offer.
	ResultsPublicationDate *DateFormat4Choice `xml:"RsltsPblctnDt,omitempty"`

	// Date/time upon which the High Court provided approval.
	CourtApprovalDate *DateFormat4Choice `xml:"CrtApprvlDt,omitempty"`

	// First possible early closing date of an offer if different from the expiry date.
	EarlyClosingDate *DateFormat4Choice `xml:"EarlyClsgDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat4Choice `xml:"ExDvddDt,omitempty"`

	// Date/time at which an index rate will be determined .
	IndexFixingDate *DateFormat4Choice `xml:"IndxFxgDt,omitempty"`

	// Date on which an interest bearing financial instrument becomes due and principal is repaid by the issuer to the investor.
	MaturityDate *DateFormat4Choice `xml:"MtrtyDt,omitempty"`

	// Date on which trading of a security is suspended as the result of an event.
	TradingSuspendedDate *DateFormat4Choice `xml:"TradgSspdDt,omitempty"`

	// Deadline by which the beneficial ownership of securities must be declared.
	CertificationDeadline *DateFormat4Choice `xml:"CertfctnDdln,omitempty"`

	// Date/time at which the securities will be redeemed (early) for payment of principal.
	RedemptionDate *DateFormat4Choice `xml:"RedDt,omitempty"`

	// Date on which instructions to register or registration details will be accepted.
	RegistrationDeadline *DateFormat4Choice `xml:"RegnDdln,omitempty"`

	// Date (and time) at which an issuer will determine the proration amount/quantity of an offer.
	ProrationDate *DateFormat4Choice `xml:"PrratnDt,omitempty"`

	// Date on until which tax breakdown instructions will be accepted.
	DeadlineForTaxBreakdownInstruction *DateFormat4Choice `xml:"DdlnForTaxBrkdwnInstr,omitempty"`

	// Date/time at which an event/offer is terminated or lapsed.
	LapsedDate *DateFormat4Choice `xml:"LpsdDt,omitempty"`

	// Last date/time by which a buying counterparty to a trade can be sure that it will have the right to participate in an event.
	GuaranteedParticipationDate *DateFormat4Choice `xml:"GrntedPrtcptnDt,omitempty"`

	// Deadline by which an entitled holder needs to advise their counterparty to a transaction of their election for a corporate action event.
	ElectionToCounterpartyDeadline *DateFormat4Choice `xml:"ElctnToCtrPtyDdln,omitempty"`

	// Date/time as from which 'special processing' can start to be used by participants for that event. Special processing is a means of marking a transaction, that would normally be traded ex or cum, as being traded cum or ex respectively, eg, a transaction dealt 'special' after the ex date would result in the buyer being eligible for the entitlement. This is typically used in the UK and Irish markets.
	SpecialExDate *DateFormat4Choice `xml:"SpclExDt,omitempty"`
}

func (c *CorporateActionDate2) AddRecordDate() *DateFormat4Choice {
	c.RecordDate = new(DateFormat4Choice)
	return c.RecordDate
}

func (c *CorporateActionDate2) AddEffectiveDate() *DateFormat4Choice {
	c.EffectiveDate = new(DateFormat4Choice)
	return c.EffectiveDate
}

func (c *CorporateActionDate2) AddCoverExpirationDate() *DateFormat4Choice {
	c.CoverExpirationDate = new(DateFormat4Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate2) AddEqualisationDate() *DateFormat4Choice {
	c.EqualisationDate = new(DateFormat4Choice)
	return c.EqualisationDate
}

func (c *CorporateActionDate2) AddMarginFixingDate() *DateFormat4Choice {
	c.MarginFixingDate = new(DateFormat4Choice)
	return c.MarginFixingDate
}

func (c *CorporateActionDate2) AddLotteryDate() *DateFormat4Choice {
	c.LotteryDate = new(DateFormat4Choice)
	return c.LotteryDate
}

func (c *CorporateActionDate2) AddProtectDate() *DateFormat4Choice {
	c.ProtectDate = new(DateFormat4Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate2) AddUnconditionalDate() *DateFormat4Choice {
	c.UnconditionalDate = new(DateFormat4Choice)
	return c.UnconditionalDate
}

func (c *CorporateActionDate2) AddWhollyUnconditionalDate() *DateFormat4Choice {
	c.WhollyUnconditionalDate = new(DateFormat4Choice)
	return c.WhollyUnconditionalDate
}

func (c *CorporateActionDate2) AddResultsPublicationDate() *DateFormat4Choice {
	c.ResultsPublicationDate = new(DateFormat4Choice)
	return c.ResultsPublicationDate
}

func (c *CorporateActionDate2) AddCourtApprovalDate() *DateFormat4Choice {
	c.CourtApprovalDate = new(DateFormat4Choice)
	return c.CourtApprovalDate
}

func (c *CorporateActionDate2) AddEarlyClosingDate() *DateFormat4Choice {
	c.EarlyClosingDate = new(DateFormat4Choice)
	return c.EarlyClosingDate
}

func (c *CorporateActionDate2) AddExDividendDate() *DateFormat4Choice {
	c.ExDividendDate = new(DateFormat4Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate2) AddIndexFixingDate() *DateFormat4Choice {
	c.IndexFixingDate = new(DateFormat4Choice)
	return c.IndexFixingDate
}

func (c *CorporateActionDate2) AddMaturityDate() *DateFormat4Choice {
	c.MaturityDate = new(DateFormat4Choice)
	return c.MaturityDate
}

func (c *CorporateActionDate2) AddTradingSuspendedDate() *DateFormat4Choice {
	c.TradingSuspendedDate = new(DateFormat4Choice)
	return c.TradingSuspendedDate
}

func (c *CorporateActionDate2) AddCertificationDeadline() *DateFormat4Choice {
	c.CertificationDeadline = new(DateFormat4Choice)
	return c.CertificationDeadline
}

func (c *CorporateActionDate2) AddRedemptionDate() *DateFormat4Choice {
	c.RedemptionDate = new(DateFormat4Choice)
	return c.RedemptionDate
}

func (c *CorporateActionDate2) AddRegistrationDeadline() *DateFormat4Choice {
	c.RegistrationDeadline = new(DateFormat4Choice)
	return c.RegistrationDeadline
}

func (c *CorporateActionDate2) AddProrationDate() *DateFormat4Choice {
	c.ProrationDate = new(DateFormat4Choice)
	return c.ProrationDate
}

func (c *CorporateActionDate2) AddDeadlineForTaxBreakdownInstruction() *DateFormat4Choice {
	c.DeadlineForTaxBreakdownInstruction = new(DateFormat4Choice)
	return c.DeadlineForTaxBreakdownInstruction
}

func (c *CorporateActionDate2) AddLapsedDate() *DateFormat4Choice {
	c.LapsedDate = new(DateFormat4Choice)
	return c.LapsedDate
}

func (c *CorporateActionDate2) AddGuaranteedParticipationDate() *DateFormat4Choice {
	c.GuaranteedParticipationDate = new(DateFormat4Choice)
	return c.GuaranteedParticipationDate
}

func (c *CorporateActionDate2) AddElectionToCounterpartyDeadline() *DateFormat4Choice {
	c.ElectionToCounterpartyDeadline = new(DateFormat4Choice)
	return c.ElectionToCounterpartyDeadline
}

func (c *CorporateActionDate2) AddSpecialExDate() *DateFormat4Choice {
	c.SpecialExDate = new(DateFormat4Choice)
	return c.SpecialExDate
}
