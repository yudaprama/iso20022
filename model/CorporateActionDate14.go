package model

// Specifies corporate action dates.
type CorporateActionDate14 struct {

	// Date/time at which the issuer announced that a corporate action event will occur.
	AnnouncementDate *DateFormat19Choice `xml:"AnncmntDt,omitempty"`

	// Deadline by which the beneficial ownership of securities must be declared.
	CertificationDeadline *DateFormat19Choice `xml:"CertfctnDdln,omitempty"`

	// Date upon which the court provided approval.
	CourtApprovalDate *DateFormat19Choice `xml:"CrtApprvlDt,omitempty"`

	// First possible early closing date of an offer if different from the expiry date.
	EarlyClosingDate *DateFormat19Choice `xml:"EarlyClsgDt,omitempty"`

	// Date/time at which an event is officially effective from the issuer's perspective.
	EffectiveDate *DateFormat19Choice `xml:"FctvDt,omitempty"`

	// Date/Time on which all or part of any holding bought in a unit trust is subject to being treated as capital rather than income. This is normally one day after the previous distribution's ex date.
	EqualisationDate *DateFormat19Choice `xml:"EqulstnDt,omitempty"`

	// Date/time at which additional information on the event will be announced, for example, exchange ratio announcement date.
	FurtherDetailedAnnouncementDate *DateFormat19Choice `xml:"FrthrDtldAnncmntDt,omitempty"`

	// Date/time at which an index rate will be determined .
	IndexFixingDate *DateFormat19Choice `xml:"IndxFxgDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat19Choice `xml:"LtryDt,omitempty"`

	// Date/time to which the maturity date of an interest bearing security is extended.
	NewMaturityDate *DateFormat19Choice `xml:"NewMtrtyDt,omitempty"`

	// Date/time on which the bondholder's or shareholder's meeting will take place.
	MeetingDate *DateFormat19Choice `xml:"MtgDt,omitempty"`

	// Date/time at which the margin rate will be determined .
	MarginFixingDate *DateFormat19Choice `xml:"MrgnFxgDt,omitempty"`

	// Date/time (and time) at which an issuer will determine the proration amount/quantity of an offer.
	ProrationDate *DateFormat19Choice `xml:"PrratnDt,omitempty"`

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat19Choice `xml:"RcrdDt,omitempty"`

	// Date/time on which instructions to register or registration details will be accepted.
	RegistrationDeadline *DateFormat19Choice `xml:"RegnDdln,omitempty"`

	// Date/time on which results are published, for example, results of an offer.
	ResultsPublicationDate *DateFormat19Choice `xml:"RsltsPblctnDt,omitempty"`

	// Deadline by which instructions must be received to split securities, for example, of physical certificates.
	DeadlineToSplit *DateFormat19Choice `xml:"DdlnToSplt,omitempty"`

	// Date/time on until which tax breakdown instructions will be accepted.
	DeadlineForTaxBreakdownInstruction *DateFormat19Choice `xml:"DdlnForTaxBrkdwnInstr,omitempty"`

	// Date/time at which trading of a security is suspended as the result of an event.
	TradingSuspendedDate *DateFormat19Choice `xml:"TradgSspdDt,omitempty"`

	// Date/time upon which the terms of the take-over become unconditional as to acceptances.
	UnconditionalDate *DateFormat19Choice `xml:"UcondlDt,omitempty"`

	// Date/time at on which all conditions, including regulatory, legal etc. pertaining to the take-over, have been met.
	WhollyUnconditionalDate *DateFormat19Choice `xml:"WhlyUcondlDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat19Choice `xml:"ExDvddDt,omitempty"`

	// Date/time at which the corporate action is legally announced by an official body, for example, publication by a governmental administration.
	OfficialAnnouncementPublicationDate *DateFormat19Choice `xml:"OffclAnncmntPblctnDt,omitempty"`

	// Date/time as from which 'special processing' can start to be used by participants for that event. Special processing is a means of marking a transaction, that would normally be traded ex or cum, as being traded cum or ex respectively,  for example, a transaction dealt 'special' after the ex date would result in the buyer being eligible for the entitlement. This is typically used in the UK and Irish markets.
	SpecialExDate *DateFormat19Choice `xml:"SpclExDt,omitempty"`

	// Last date/time by which a buying counterparty to a trade can be sure that it will have the right to participate in an event.
	GuaranteedParticipationDate *DateFormat19Choice `xml:"GrntedPrtcptnDt,omitempty"`

	// Deadline by which an entitled holder needs to advise their counterparty to a transaction of their election for a corporate action event.
	ElectionToCounterpartyDeadline *DateFormat19Choice `xml:"ElctnToCtrPtyDdln,omitempty"`

	// Date/time at which an event/offer is terminated or lapsed.
	LapsedDate *DateFormat19Choice `xml:"LpsdDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat19Choice `xml:"PmtDt,omitempty"`

	// Date/Time by which the account owner must instruct directly another party, for example to provide documentation to an issuer agent.
	ThirdPartyDeadline *DateFormat19Choice `xml:"ThrdPtyDdln,omitempty"`

	// Date/Time set by the issuer agent as a first early deadline by which the account owner must instruct directly another party, possibly giving the holder eligibility to incentives. For example, to provide documentation to an issuer agent.
	EarlyThirdPartyDeadline *DateFormat19Choice `xml:"EarlyThrdPtyDdln,omitempty"`

	// Date by which the depository stops monitoring activities of the event, for instance, accounting and tracking activities for due bills end.
	MarketClaimTrackingEndDate *DateFormat19Choice `xml:"MktClmTrckgEndDt,omitempty"`

	// Last day an investor can become a lead plaintiff.
	LeadPlaintiffDeadline *DateFormat19Choice `xml:"LeadPlntffDdln,omitempty"`
}

func (c *CorporateActionDate14) AddAnnouncementDate() *DateFormat19Choice {
	c.AnnouncementDate = new(DateFormat19Choice)
	return c.AnnouncementDate
}

func (c *CorporateActionDate14) AddCertificationDeadline() *DateFormat19Choice {
	c.CertificationDeadline = new(DateFormat19Choice)
	return c.CertificationDeadline
}

func (c *CorporateActionDate14) AddCourtApprovalDate() *DateFormat19Choice {
	c.CourtApprovalDate = new(DateFormat19Choice)
	return c.CourtApprovalDate
}

func (c *CorporateActionDate14) AddEarlyClosingDate() *DateFormat19Choice {
	c.EarlyClosingDate = new(DateFormat19Choice)
	return c.EarlyClosingDate
}

func (c *CorporateActionDate14) AddEffectiveDate() *DateFormat19Choice {
	c.EffectiveDate = new(DateFormat19Choice)
	return c.EffectiveDate
}

func (c *CorporateActionDate14) AddEqualisationDate() *DateFormat19Choice {
	c.EqualisationDate = new(DateFormat19Choice)
	return c.EqualisationDate
}

func (c *CorporateActionDate14) AddFurtherDetailedAnnouncementDate() *DateFormat19Choice {
	c.FurtherDetailedAnnouncementDate = new(DateFormat19Choice)
	return c.FurtherDetailedAnnouncementDate
}

func (c *CorporateActionDate14) AddIndexFixingDate() *DateFormat19Choice {
	c.IndexFixingDate = new(DateFormat19Choice)
	return c.IndexFixingDate
}

func (c *CorporateActionDate14) AddLotteryDate() *DateFormat19Choice {
	c.LotteryDate = new(DateFormat19Choice)
	return c.LotteryDate
}

func (c *CorporateActionDate14) AddNewMaturityDate() *DateFormat19Choice {
	c.NewMaturityDate = new(DateFormat19Choice)
	return c.NewMaturityDate
}

func (c *CorporateActionDate14) AddMeetingDate() *DateFormat19Choice {
	c.MeetingDate = new(DateFormat19Choice)
	return c.MeetingDate
}

func (c *CorporateActionDate14) AddMarginFixingDate() *DateFormat19Choice {
	c.MarginFixingDate = new(DateFormat19Choice)
	return c.MarginFixingDate
}

func (c *CorporateActionDate14) AddProrationDate() *DateFormat19Choice {
	c.ProrationDate = new(DateFormat19Choice)
	return c.ProrationDate
}

func (c *CorporateActionDate14) AddRecordDate() *DateFormat19Choice {
	c.RecordDate = new(DateFormat19Choice)
	return c.RecordDate
}

func (c *CorporateActionDate14) AddRegistrationDeadline() *DateFormat19Choice {
	c.RegistrationDeadline = new(DateFormat19Choice)
	return c.RegistrationDeadline
}

func (c *CorporateActionDate14) AddResultsPublicationDate() *DateFormat19Choice {
	c.ResultsPublicationDate = new(DateFormat19Choice)
	return c.ResultsPublicationDate
}

func (c *CorporateActionDate14) AddDeadlineToSplit() *DateFormat19Choice {
	c.DeadlineToSplit = new(DateFormat19Choice)
	return c.DeadlineToSplit
}

func (c *CorporateActionDate14) AddDeadlineForTaxBreakdownInstruction() *DateFormat19Choice {
	c.DeadlineForTaxBreakdownInstruction = new(DateFormat19Choice)
	return c.DeadlineForTaxBreakdownInstruction
}

func (c *CorporateActionDate14) AddTradingSuspendedDate() *DateFormat19Choice {
	c.TradingSuspendedDate = new(DateFormat19Choice)
	return c.TradingSuspendedDate
}

func (c *CorporateActionDate14) AddUnconditionalDate() *DateFormat19Choice {
	c.UnconditionalDate = new(DateFormat19Choice)
	return c.UnconditionalDate
}

func (c *CorporateActionDate14) AddWhollyUnconditionalDate() *DateFormat19Choice {
	c.WhollyUnconditionalDate = new(DateFormat19Choice)
	return c.WhollyUnconditionalDate
}

func (c *CorporateActionDate14) AddExDividendDate() *DateFormat19Choice {
	c.ExDividendDate = new(DateFormat19Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate14) AddOfficialAnnouncementPublicationDate() *DateFormat19Choice {
	c.OfficialAnnouncementPublicationDate = new(DateFormat19Choice)
	return c.OfficialAnnouncementPublicationDate
}

func (c *CorporateActionDate14) AddSpecialExDate() *DateFormat19Choice {
	c.SpecialExDate = new(DateFormat19Choice)
	return c.SpecialExDate
}

func (c *CorporateActionDate14) AddGuaranteedParticipationDate() *DateFormat19Choice {
	c.GuaranteedParticipationDate = new(DateFormat19Choice)
	return c.GuaranteedParticipationDate
}

func (c *CorporateActionDate14) AddElectionToCounterpartyDeadline() *DateFormat19Choice {
	c.ElectionToCounterpartyDeadline = new(DateFormat19Choice)
	return c.ElectionToCounterpartyDeadline
}

func (c *CorporateActionDate14) AddLapsedDate() *DateFormat19Choice {
	c.LapsedDate = new(DateFormat19Choice)
	return c.LapsedDate
}

func (c *CorporateActionDate14) AddPaymentDate() *DateFormat19Choice {
	c.PaymentDate = new(DateFormat19Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate14) AddThirdPartyDeadline() *DateFormat19Choice {
	c.ThirdPartyDeadline = new(DateFormat19Choice)
	return c.ThirdPartyDeadline
}

func (c *CorporateActionDate14) AddEarlyThirdPartyDeadline() *DateFormat19Choice {
	c.EarlyThirdPartyDeadline = new(DateFormat19Choice)
	return c.EarlyThirdPartyDeadline
}

func (c *CorporateActionDate14) AddMarketClaimTrackingEndDate() *DateFormat19Choice {
	c.MarketClaimTrackingEndDate = new(DateFormat19Choice)
	return c.MarketClaimTrackingEndDate
}

func (c *CorporateActionDate14) AddLeadPlaintiffDeadline() *DateFormat19Choice {
	c.LeadPlaintiffDeadline = new(DateFormat19Choice)
	return c.LeadPlaintiffDeadline
}
