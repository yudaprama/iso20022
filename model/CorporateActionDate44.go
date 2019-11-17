package model

// Specifies corporate action dates.
type CorporateActionDate44 struct {

	// Date/time at which the issuer announced that a corporate action event will occur.
	AnnouncementDate *DateFormat31Choice `xml:"AnncmntDt,omitempty"`

	// Deadline by which the beneficial ownership of securities must be declared.
	CertificationDeadline *DateFormat31Choice `xml:"CertfctnDdln,omitempty"`

	// Date upon which the court provided approval.
	CourtApprovalDate *DateFormat31Choice `xml:"CrtApprvlDt,omitempty"`

	// First possible early closing date of an offer if different from the expiry date.
	EarlyClosingDate *DateFormat31Choice `xml:"EarlyClsgDt,omitempty"`

	// Date/time at which an event is officially effective from the issuer's perspective.
	EffectiveDate *DateFormat31Choice `xml:"FctvDt,omitempty"`

	// Date/Time on which all or part of any holding bought in a unit trust is subject to being treated as capital rather than income. This is normally one day after the previous distribution's ex date.
	EqualisationDate *DateFormat31Choice `xml:"EqulstnDt,omitempty"`

	// Date/time at which additional information on the event will be announced, for example, exchange ratio announcement date.
	FurtherDetailedAnnouncementDate *DateFormat31Choice `xml:"FrthrDtldAnncmntDt,omitempty"`

	// Date/time at which an index / rate / price / value will be determined.
	FixingDate *DateFormat31Choice `xml:"FxgDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat31Choice `xml:"LtryDt,omitempty"`

	// Date/time to which the maturity date of an interest bearing security is extended.
	NewMaturityDate *DateFormat31Choice `xml:"NewMtrtyDt,omitempty"`

	// Date/time on which the bondholder's or shareholder's meeting will take place.
	MeetingDate *DateFormat31Choice `xml:"MtgDt,omitempty"`

	// Date/time at which the margin rate will be determined .
	MarginFixingDate *DateFormat31Choice `xml:"MrgnFxgDt,omitempty"`

	// Date/time (and time) at which an issuer will determine the proration amount/quantity of an offer.
	ProrationDate *DateFormat31Choice `xml:"PrratnDt,omitempty"`

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat31Choice `xml:"RcrdDt,omitempty"`

	// Date/time on which instructions to register or registration details will be accepted.
	RegistrationDeadline *DateFormat31Choice `xml:"RegnDdln,omitempty"`

	// Date/time on which results are published, for example, results of an offer.
	ResultsPublicationDate *DateFormat31Choice `xml:"RsltsPblctnDt,omitempty"`

	// Deadline by which instructions must be received to split securities, for example, of physical certificates.
	DeadlineToSplit *DateFormat31Choice `xml:"DdlnToSplt,omitempty"`

	// Date/time on until which tax breakdown instructions will be accepted.
	DeadlineForTaxBreakdownInstruction *DateFormat31Choice `xml:"DdlnForTaxBrkdwnInstr,omitempty"`

	// Date/time at which trading of a security is suspended as the result of an event.
	TradingSuspendedDate *DateFormat31Choice `xml:"TradgSspdDt,omitempty"`

	// Date/time upon which the terms of the take-over become unconditional as to acceptances.
	UnconditionalDate *DateFormat31Choice `xml:"UcondlDt,omitempty"`

	// Date/time at on which all conditions, including regulatory, legal etc. pertaining to the take-over, have been met.
	WhollyUnconditionalDate *DateFormat31Choice `xml:"WhlyUcondlDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat31Choice `xml:"ExDvddDt,omitempty"`

	// Date/time at which the corporate action is legally announced by an official body, for example, publication by a governmental administration.
	OfficialAnnouncementPublicationDate *DateFormat31Choice `xml:"OffclAnncmntPblctnDt,omitempty"`

	// Date/time as from which 'special processing' can start to be used by participants for that event. Special processing is a means of marking a transaction, that would normally be traded ex or cum, as being traded cum or ex respectively,  for example, a transaction dealt 'special' after the ex date would result in the buyer being eligible for the entitlement. This is typically used in the UK and Irish markets.
	SpecialExDate *DateFormat31Choice `xml:"SpclExDt,omitempty"`

	// Last date/time by which a buying counterparty to a trade can be sure that it will have the right to participate in an event.
	GuaranteedParticipationDate *DateFormat31Choice `xml:"GrntedPrtcptnDt,omitempty"`

	// Deadline by which an entitled holder needs to advise their counterparty to a transaction of their election for a corporate action event, also known as Buyer Protection Deadline.
	ElectionToCounterpartyMarketDeadline *DateFormat31Choice `xml:"ElctnToCtrPtyMktDdln,omitempty"`

	// Date/time the account servicer has set as the deadline to respond, with instructions, prior to the election to counterparty market deadline
	ElectionToCounterpartyResponseDeadline *DateFormat31Choice `xml:"ElctnToCtrPtyRspnDdln,omitempty"`

	// Date/time at which an event/offer is terminated or lapsed.
	LapsedDate *DateFormat31Choice `xml:"LpsdDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat31Choice `xml:"PmtDt,omitempty"`

	// Date/Time by which the account owner must instruct directly another party, for example to provide documentation to an issuer agent.
	ThirdPartyDeadline *DateFormat31Choice `xml:"ThrdPtyDdln,omitempty"`

	// Date/Time set by the issuer agent as a first early deadline by which the account owner must instruct directly another party, possibly giving the holder eligibility to incentives. For example, to provide documentation to an issuer agent.
	EarlyThirdPartyDeadline *DateFormat31Choice `xml:"EarlyThrdPtyDdln,omitempty"`

	// Date by which the depository stops monitoring activities of the event, for instance, accounting and tracking activities for due bills end.
	MarketClaimTrackingEndDate *DateFormat31Choice `xml:"MktClmTrckgEndDt,omitempty"`

	// Last day an investor can become a lead plaintiff.
	LeadPlaintiffDeadline *DateFormat31Choice `xml:"LeadPlntffDdln,omitempty"`

	// Date on which the action was filed at the applicable court.
	FilingDate *DateFormat30Choice `xml:"FilgDt,omitempty"`

	// Date for the hearing between the plaintiff and defendant, as set by the court.
	HearingDate *DateFormat30Choice `xml:"HrgDt,omitempty"`
}

func (c *CorporateActionDate44) AddAnnouncementDate() *DateFormat31Choice {
	c.AnnouncementDate = new(DateFormat31Choice)
	return c.AnnouncementDate
}

func (c *CorporateActionDate44) AddCertificationDeadline() *DateFormat31Choice {
	c.CertificationDeadline = new(DateFormat31Choice)
	return c.CertificationDeadline
}

func (c *CorporateActionDate44) AddCourtApprovalDate() *DateFormat31Choice {
	c.CourtApprovalDate = new(DateFormat31Choice)
	return c.CourtApprovalDate
}

func (c *CorporateActionDate44) AddEarlyClosingDate() *DateFormat31Choice {
	c.EarlyClosingDate = new(DateFormat31Choice)
	return c.EarlyClosingDate
}

func (c *CorporateActionDate44) AddEffectiveDate() *DateFormat31Choice {
	c.EffectiveDate = new(DateFormat31Choice)
	return c.EffectiveDate
}

func (c *CorporateActionDate44) AddEqualisationDate() *DateFormat31Choice {
	c.EqualisationDate = new(DateFormat31Choice)
	return c.EqualisationDate
}

func (c *CorporateActionDate44) AddFurtherDetailedAnnouncementDate() *DateFormat31Choice {
	c.FurtherDetailedAnnouncementDate = new(DateFormat31Choice)
	return c.FurtherDetailedAnnouncementDate
}

func (c *CorporateActionDate44) AddFixingDate() *DateFormat31Choice {
	c.FixingDate = new(DateFormat31Choice)
	return c.FixingDate
}

func (c *CorporateActionDate44) AddLotteryDate() *DateFormat31Choice {
	c.LotteryDate = new(DateFormat31Choice)
	return c.LotteryDate
}

func (c *CorporateActionDate44) AddNewMaturityDate() *DateFormat31Choice {
	c.NewMaturityDate = new(DateFormat31Choice)
	return c.NewMaturityDate
}

func (c *CorporateActionDate44) AddMeetingDate() *DateFormat31Choice {
	c.MeetingDate = new(DateFormat31Choice)
	return c.MeetingDate
}

func (c *CorporateActionDate44) AddMarginFixingDate() *DateFormat31Choice {
	c.MarginFixingDate = new(DateFormat31Choice)
	return c.MarginFixingDate
}

func (c *CorporateActionDate44) AddProrationDate() *DateFormat31Choice {
	c.ProrationDate = new(DateFormat31Choice)
	return c.ProrationDate
}

func (c *CorporateActionDate44) AddRecordDate() *DateFormat31Choice {
	c.RecordDate = new(DateFormat31Choice)
	return c.RecordDate
}

func (c *CorporateActionDate44) AddRegistrationDeadline() *DateFormat31Choice {
	c.RegistrationDeadline = new(DateFormat31Choice)
	return c.RegistrationDeadline
}

func (c *CorporateActionDate44) AddResultsPublicationDate() *DateFormat31Choice {
	c.ResultsPublicationDate = new(DateFormat31Choice)
	return c.ResultsPublicationDate
}

func (c *CorporateActionDate44) AddDeadlineToSplit() *DateFormat31Choice {
	c.DeadlineToSplit = new(DateFormat31Choice)
	return c.DeadlineToSplit
}

func (c *CorporateActionDate44) AddDeadlineForTaxBreakdownInstruction() *DateFormat31Choice {
	c.DeadlineForTaxBreakdownInstruction = new(DateFormat31Choice)
	return c.DeadlineForTaxBreakdownInstruction
}

func (c *CorporateActionDate44) AddTradingSuspendedDate() *DateFormat31Choice {
	c.TradingSuspendedDate = new(DateFormat31Choice)
	return c.TradingSuspendedDate
}

func (c *CorporateActionDate44) AddUnconditionalDate() *DateFormat31Choice {
	c.UnconditionalDate = new(DateFormat31Choice)
	return c.UnconditionalDate
}

func (c *CorporateActionDate44) AddWhollyUnconditionalDate() *DateFormat31Choice {
	c.WhollyUnconditionalDate = new(DateFormat31Choice)
	return c.WhollyUnconditionalDate
}

func (c *CorporateActionDate44) AddExDividendDate() *DateFormat31Choice {
	c.ExDividendDate = new(DateFormat31Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate44) AddOfficialAnnouncementPublicationDate() *DateFormat31Choice {
	c.OfficialAnnouncementPublicationDate = new(DateFormat31Choice)
	return c.OfficialAnnouncementPublicationDate
}

func (c *CorporateActionDate44) AddSpecialExDate() *DateFormat31Choice {
	c.SpecialExDate = new(DateFormat31Choice)
	return c.SpecialExDate
}

func (c *CorporateActionDate44) AddGuaranteedParticipationDate() *DateFormat31Choice {
	c.GuaranteedParticipationDate = new(DateFormat31Choice)
	return c.GuaranteedParticipationDate
}

func (c *CorporateActionDate44) AddElectionToCounterpartyMarketDeadline() *DateFormat31Choice {
	c.ElectionToCounterpartyMarketDeadline = new(DateFormat31Choice)
	return c.ElectionToCounterpartyMarketDeadline
}

func (c *CorporateActionDate44) AddElectionToCounterpartyResponseDeadline() *DateFormat31Choice {
	c.ElectionToCounterpartyResponseDeadline = new(DateFormat31Choice)
	return c.ElectionToCounterpartyResponseDeadline
}

func (c *CorporateActionDate44) AddLapsedDate() *DateFormat31Choice {
	c.LapsedDate = new(DateFormat31Choice)
	return c.LapsedDate
}

func (c *CorporateActionDate44) AddPaymentDate() *DateFormat31Choice {
	c.PaymentDate = new(DateFormat31Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate44) AddThirdPartyDeadline() *DateFormat31Choice {
	c.ThirdPartyDeadline = new(DateFormat31Choice)
	return c.ThirdPartyDeadline
}

func (c *CorporateActionDate44) AddEarlyThirdPartyDeadline() *DateFormat31Choice {
	c.EarlyThirdPartyDeadline = new(DateFormat31Choice)
	return c.EarlyThirdPartyDeadline
}

func (c *CorporateActionDate44) AddMarketClaimTrackingEndDate() *DateFormat31Choice {
	c.MarketClaimTrackingEndDate = new(DateFormat31Choice)
	return c.MarketClaimTrackingEndDate
}

func (c *CorporateActionDate44) AddLeadPlaintiffDeadline() *DateFormat31Choice {
	c.LeadPlaintiffDeadline = new(DateFormat31Choice)
	return c.LeadPlaintiffDeadline
}

func (c *CorporateActionDate44) AddFilingDate() *DateFormat30Choice {
	c.FilingDate = new(DateFormat30Choice)
	return c.FilingDate
}

func (c *CorporateActionDate44) AddHearingDate() *DateFormat30Choice {
	c.HearingDate = new(DateFormat30Choice)
	return c.HearingDate
}
