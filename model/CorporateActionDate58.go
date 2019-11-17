package model

// Specifies corporate action dates.
type CorporateActionDate58 struct {

	// Date/time at which the issuer announced that a corporate action event will occur.
	AnnouncementDate *DateFormat34Choice `xml:"AnncmntDt,omitempty"`

	// Deadline by which the beneficial ownership of securities must be declared.
	CertificationDeadline *DateFormat34Choice `xml:"CertfctnDdln,omitempty"`

	// Date upon which the court provided approval.
	CourtApprovalDate *DateFormat34Choice `xml:"CrtApprvlDt,omitempty"`

	// First possible early closing date of an offer if different from the expiry date.
	EarlyClosingDate *DateFormat34Choice `xml:"EarlyClsgDt,omitempty"`

	// Date/time at which an event is officially effective from the issuer's perspective.
	EffectiveDate *DateFormat34Choice `xml:"FctvDt,omitempty"`

	// Date/Time on which all or part of any holding bought in a unit trust is subject to being treated as capital rather than income. This is normally one day after the previous distribution's ex date.
	EqualisationDate *DateFormat34Choice `xml:"EqulstnDt,omitempty"`

	// Date/time at which additional information on the event will be announced, for example, exchange ratio announcement date.
	FurtherDetailedAnnouncementDate *DateFormat34Choice `xml:"FrthrDtldAnncmntDt,omitempty"`

	// Date/time at which an index / rate / price / value will be determined.
	FixingDate *DateFormat34Choice `xml:"FxgDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat34Choice `xml:"LtryDt,omitempty"`

	// Date/time to which the maturity date of an interest bearing security is extended.
	NewMaturityDate *DateFormat34Choice `xml:"NewMtrtyDt,omitempty"`

	// Date/time on which the bondholder's or shareholder's meeting will take place.
	MeetingDate *DateFormat34Choice `xml:"MtgDt,omitempty"`

	// Date/time at which the margin rate will be determined .
	MarginFixingDate *DateFormat34Choice `xml:"MrgnFxgDt,omitempty"`

	// Date/time (and time) at which an issuer will determine the proration amount/quantity of an offer.
	ProrationDate *DateFormat34Choice `xml:"PrratnDt,omitempty"`

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat34Choice `xml:"RcrdDt,omitempty"`

	// Date/time on which instructions to register or registration details will be accepted.
	RegistrationDeadline *DateFormat34Choice `xml:"RegnDdln,omitempty"`

	// Date/time on which results are published, for example, results of an offer.
	ResultsPublicationDate *DateFormat34Choice `xml:"RsltsPblctnDt,omitempty"`

	// Deadline by which instructions must be received to split securities, for example, of physical certificates.
	DeadlineToSplit *DateFormat34Choice `xml:"DdlnToSplt,omitempty"`

	// Date/time on until which tax breakdown instructions will be accepted.
	DeadlineForTaxBreakdownInstruction *DateFormat34Choice `xml:"DdlnForTaxBrkdwnInstr,omitempty"`

	// Date/time at which trading of a security is suspended as the result of an event.
	TradingSuspendedDate *DateFormat34Choice `xml:"TradgSspdDt,omitempty"`

	// Date/time upon which the terms of the take-over become unconditional as to acceptances.
	UnconditionalDate *DateFormat34Choice `xml:"UcondlDt,omitempty"`

	// Date/time at on which all conditions, including regulatory, legal etc. pertaining to the take-over, have been met.
	WhollyUnconditionalDate *DateFormat34Choice `xml:"WhlyUcondlDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat34Choice `xml:"ExDvddDt,omitempty"`

	// Date/time at which the corporate action is legally announced by an official body, for example, publication by a governmental administration.
	OfficialAnnouncementPublicationDate *DateFormat34Choice `xml:"OffclAnncmntPblctnDt,omitempty"`

	// Date/time as from which 'special processing' can start to be used by participants for that event. Special processing is a means of marking a transaction, that would normally be traded ex or cum, as being traded cum or ex respectively,  for example, a transaction dealt 'special' after the ex date would result in the buyer being eligible for the entitlement. This is typically used in the UK and Irish markets.
	SpecialExDate *DateFormat34Choice `xml:"SpclExDt,omitempty"`

	// Last date/time by which a buying counterparty to a trade can be sure that it will have the right to participate in an event.
	GuaranteedParticipationDate *DateFormat34Choice `xml:"GrntedPrtcptnDt,omitempty"`

	// Deadline by which an entitled holder needs to advise their counterparty to a transaction of their election for a corporate action event, also known as Buyer Protection Deadline.
	ElectionToCounterpartyMarketDeadline *DateFormat34Choice `xml:"ElctnToCtrPtyMktDdln,omitempty"`

	// Date/time the account servicer has set as the deadline to respond, with instructions, prior to the election to counterparty market deadline
	ElectionToCounterpartyResponseDeadline *DateFormat34Choice `xml:"ElctnToCtrPtyRspnDdln,omitempty"`

	// Date/time at which an event/offer is terminated or lapsed.
	LapsedDate *DateFormat34Choice `xml:"LpsdDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat34Choice `xml:"PmtDt,omitempty"`

	// Date/Time by which the account owner must instruct directly another party, for example to provide documentation to an issuer agent.
	ThirdPartyDeadline *DateFormat34Choice `xml:"ThrdPtyDdln,omitempty"`

	// Date/Time set by the issuer agent as a first early deadline by which the account owner must instruct directly another party, possibly giving the holder eligibility to incentives. For example, to provide documentation to an issuer agent.
	EarlyThirdPartyDeadline *DateFormat34Choice `xml:"EarlyThrdPtyDdln,omitempty"`

	// Date by which the depository stops monitoring activities of the event, for instance, accounting and tracking activities for due bills end.
	MarketClaimTrackingEndDate *DateFormat34Choice `xml:"MktClmTrckgEndDt,omitempty"`

	// Last day an investor can become a lead plaintiff.
	LeadPlaintiffDeadline *DateFormat34Choice `xml:"LeadPlntffDdln,omitempty"`

	// Date on which the action was filed at the applicable court.
	FilingDate *DateFormat41Choice `xml:"FilgDt,omitempty"`

	// Date for the hearing between the plaintiff and defendant, as set by the court.
	HearingDate *DateFormat41Choice `xml:"HrgDt,omitempty"`
}

func (c *CorporateActionDate58) AddAnnouncementDate() *DateFormat34Choice {
	c.AnnouncementDate = new(DateFormat34Choice)
	return c.AnnouncementDate
}

func (c *CorporateActionDate58) AddCertificationDeadline() *DateFormat34Choice {
	c.CertificationDeadline = new(DateFormat34Choice)
	return c.CertificationDeadline
}

func (c *CorporateActionDate58) AddCourtApprovalDate() *DateFormat34Choice {
	c.CourtApprovalDate = new(DateFormat34Choice)
	return c.CourtApprovalDate
}

func (c *CorporateActionDate58) AddEarlyClosingDate() *DateFormat34Choice {
	c.EarlyClosingDate = new(DateFormat34Choice)
	return c.EarlyClosingDate
}

func (c *CorporateActionDate58) AddEffectiveDate() *DateFormat34Choice {
	c.EffectiveDate = new(DateFormat34Choice)
	return c.EffectiveDate
}

func (c *CorporateActionDate58) AddEqualisationDate() *DateFormat34Choice {
	c.EqualisationDate = new(DateFormat34Choice)
	return c.EqualisationDate
}

func (c *CorporateActionDate58) AddFurtherDetailedAnnouncementDate() *DateFormat34Choice {
	c.FurtherDetailedAnnouncementDate = new(DateFormat34Choice)
	return c.FurtherDetailedAnnouncementDate
}

func (c *CorporateActionDate58) AddFixingDate() *DateFormat34Choice {
	c.FixingDate = new(DateFormat34Choice)
	return c.FixingDate
}

func (c *CorporateActionDate58) AddLotteryDate() *DateFormat34Choice {
	c.LotteryDate = new(DateFormat34Choice)
	return c.LotteryDate
}

func (c *CorporateActionDate58) AddNewMaturityDate() *DateFormat34Choice {
	c.NewMaturityDate = new(DateFormat34Choice)
	return c.NewMaturityDate
}

func (c *CorporateActionDate58) AddMeetingDate() *DateFormat34Choice {
	c.MeetingDate = new(DateFormat34Choice)
	return c.MeetingDate
}

func (c *CorporateActionDate58) AddMarginFixingDate() *DateFormat34Choice {
	c.MarginFixingDate = new(DateFormat34Choice)
	return c.MarginFixingDate
}

func (c *CorporateActionDate58) AddProrationDate() *DateFormat34Choice {
	c.ProrationDate = new(DateFormat34Choice)
	return c.ProrationDate
}

func (c *CorporateActionDate58) AddRecordDate() *DateFormat34Choice {
	c.RecordDate = new(DateFormat34Choice)
	return c.RecordDate
}

func (c *CorporateActionDate58) AddRegistrationDeadline() *DateFormat34Choice {
	c.RegistrationDeadline = new(DateFormat34Choice)
	return c.RegistrationDeadline
}

func (c *CorporateActionDate58) AddResultsPublicationDate() *DateFormat34Choice {
	c.ResultsPublicationDate = new(DateFormat34Choice)
	return c.ResultsPublicationDate
}

func (c *CorporateActionDate58) AddDeadlineToSplit() *DateFormat34Choice {
	c.DeadlineToSplit = new(DateFormat34Choice)
	return c.DeadlineToSplit
}

func (c *CorporateActionDate58) AddDeadlineForTaxBreakdownInstruction() *DateFormat34Choice {
	c.DeadlineForTaxBreakdownInstruction = new(DateFormat34Choice)
	return c.DeadlineForTaxBreakdownInstruction
}

func (c *CorporateActionDate58) AddTradingSuspendedDate() *DateFormat34Choice {
	c.TradingSuspendedDate = new(DateFormat34Choice)
	return c.TradingSuspendedDate
}

func (c *CorporateActionDate58) AddUnconditionalDate() *DateFormat34Choice {
	c.UnconditionalDate = new(DateFormat34Choice)
	return c.UnconditionalDate
}

func (c *CorporateActionDate58) AddWhollyUnconditionalDate() *DateFormat34Choice {
	c.WhollyUnconditionalDate = new(DateFormat34Choice)
	return c.WhollyUnconditionalDate
}

func (c *CorporateActionDate58) AddExDividendDate() *DateFormat34Choice {
	c.ExDividendDate = new(DateFormat34Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate58) AddOfficialAnnouncementPublicationDate() *DateFormat34Choice {
	c.OfficialAnnouncementPublicationDate = new(DateFormat34Choice)
	return c.OfficialAnnouncementPublicationDate
}

func (c *CorporateActionDate58) AddSpecialExDate() *DateFormat34Choice {
	c.SpecialExDate = new(DateFormat34Choice)
	return c.SpecialExDate
}

func (c *CorporateActionDate58) AddGuaranteedParticipationDate() *DateFormat34Choice {
	c.GuaranteedParticipationDate = new(DateFormat34Choice)
	return c.GuaranteedParticipationDate
}

func (c *CorporateActionDate58) AddElectionToCounterpartyMarketDeadline() *DateFormat34Choice {
	c.ElectionToCounterpartyMarketDeadline = new(DateFormat34Choice)
	return c.ElectionToCounterpartyMarketDeadline
}

func (c *CorporateActionDate58) AddElectionToCounterpartyResponseDeadline() *DateFormat34Choice {
	c.ElectionToCounterpartyResponseDeadline = new(DateFormat34Choice)
	return c.ElectionToCounterpartyResponseDeadline
}

func (c *CorporateActionDate58) AddLapsedDate() *DateFormat34Choice {
	c.LapsedDate = new(DateFormat34Choice)
	return c.LapsedDate
}

func (c *CorporateActionDate58) AddPaymentDate() *DateFormat34Choice {
	c.PaymentDate = new(DateFormat34Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate58) AddThirdPartyDeadline() *DateFormat34Choice {
	c.ThirdPartyDeadline = new(DateFormat34Choice)
	return c.ThirdPartyDeadline
}

func (c *CorporateActionDate58) AddEarlyThirdPartyDeadline() *DateFormat34Choice {
	c.EarlyThirdPartyDeadline = new(DateFormat34Choice)
	return c.EarlyThirdPartyDeadline
}

func (c *CorporateActionDate58) AddMarketClaimTrackingEndDate() *DateFormat34Choice {
	c.MarketClaimTrackingEndDate = new(DateFormat34Choice)
	return c.MarketClaimTrackingEndDate
}

func (c *CorporateActionDate58) AddLeadPlaintiffDeadline() *DateFormat34Choice {
	c.LeadPlaintiffDeadline = new(DateFormat34Choice)
	return c.LeadPlaintiffDeadline
}

func (c *CorporateActionDate58) AddFilingDate() *DateFormat41Choice {
	c.FilingDate = new(DateFormat41Choice)
	return c.FilingDate
}

func (c *CorporateActionDate58) AddHearingDate() *DateFormat41Choice {
	c.HearingDate = new(DateFormat41Choice)
	return c.HearingDate
}
