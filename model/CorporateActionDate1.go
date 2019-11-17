package model

// Specifies corporate action dates.
type CorporateActionDate1 struct {

	// Date/time at which the issuer announced that a corporate action event will occur.
	AnnouncementDate *DateFormat6Choice `xml:"AnncmntDt,omitempty"`

	// Deadline by which the beneficial ownership of securities must be declared.
	CertificationDeadline *DateFormat6Choice `xml:"CertfctnDdln,omitempty"`

	// Date upon which the court provided approval.
	CourtApprovalDate *DateFormat6Choice `xml:"CrtApprvlDt,omitempty"`

	// First possible early closing date of an offer if different from the expiry date.
	EarlyClosingDate *DateFormat6Choice `xml:"EarlyClsgDt,omitempty"`

	// Date/time at which an event is officially effective from the issuer's perspective.
	EffectiveDate *DateFormat6Choice `xml:"FctvDt,omitempty"`

	// Date/Time on which all or part of any holding bought in a unit trust is subject to being treated as capital rather than income. This is normally one day after the previous distribution's ex date.
	EqualisationDate *DateFormat6Choice `xml:"EqulstnDt,omitempty"`

	// Date/time at which additional information on the event will be announced, for example, exchange ratio announcement date.
	FurtherDetailedAnnouncementDate *DateFormat6Choice `xml:"FrthrDtldAnncmntDt,omitempty"`

	// Date/time at which an index rate will be determined .
	IndexFixingDate *DateFormat6Choice `xml:"IndxFxgDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat6Choice `xml:"LtryDt,omitempty"`

	// Date/time upon on which an interest bearing financial instrument becomes due and principal is repaid by the issuer to the investor.
	MaturityDate *DateFormat6Choice `xml:"MtrtyDt,omitempty"`

	// Date/time on which the bondholder's or shareholder's meeting will take place.
	MeetingDate *DateFormat6Choice `xml:"MtgDt,omitempty"`

	// Date/time at which the margin rate will be determined .
	MarginFixingDate *DateFormat6Choice `xml:"MrgnFxgDt,omitempty"`

	// Date/time (and time) at which an issuer will determine the proration amount/quantity of an offer.
	ProrationDate *DateFormat6Choice `xml:"PrratnDt,omitempty"`

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat6Choice `xml:"RcrdDt,omitempty"`

	// Date/time on which instructions to register or registration details will be accepted.
	RegistrationDeadline *DateFormat6Choice `xml:"RegnDdln,omitempty"`

	// Date/time on which results are published, for example, results of an offer.
	ResultsPublicationDate *DateFormat6Choice `xml:"RsltsPblctnDt,omitempty"`

	// Deadline by which instructions must be received to split securities, for example, of physical certificates.
	DeadlineToSplit *DateFormat6Choice `xml:"DdlnToSplt,omitempty"`

	// Date/time on until which tax breakdown instructions will be accepted.
	DeadlineForTaxBreakdownInstruction *DateFormat6Choice `xml:"DdlnForTaxBrkdwnInstr,omitempty"`

	// Date/time at which trading of a security is suspended as the result of an event.
	TradingSuspendedDate *DateFormat6Choice `xml:"TradgSspdDt,omitempty"`

	// Date/time upon which the terms of the take-over become unconditional as to acceptances.
	UnconditionalDate *DateFormat6Choice `xml:"UcondlDt,omitempty"`

	// Date/time at on which all conditions, including regulatory, legal etc. pertaining to the take-over, have been met.
	WhollyUnconditionalDate *DateFormat6Choice `xml:"WhlyUcondlDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat6Choice `xml:"ExDvddDt,omitempty"`

	// Date/time at which the corporate action is legally announced by an official body, for example, publication by a governmental administration.
	OfficialAnnouncementPublicationDate *DateFormat6Choice `xml:"OffclAnncmntPblctnDt,omitempty"`

	// Date/time as from which 'special processing' can start to be used by participants for that event. Special processing is a means of marking a transaction, that would normally be traded ex or cum, as being traded cum or ex respectively,  for example, a transaction dealt 'special' after the ex date would result in the buyer being eligible for the entitlement. This is typically used in the UK and Irish markets.
	SpecialExDate *DateFormat6Choice `xml:"SpclExDt,omitempty"`

	// Last date/time by which a buying counterparty to a trade can be sure that it will have the right to participate in an event.
	GuaranteedParticipationDate *DateFormat6Choice `xml:"GrntedPrtcptnDt,omitempty"`

	// Deadline by which an entitled holder needs to advise their counterparty to a transaction of their election for a corporate action event.
	ElectionToCounterpartyDeadline *DateFormat6Choice `xml:"ElctnToCtrPtyDdln,omitempty"`

	// Date/time at which an event/offer is terminated or lapsed.
	LapsedDate *DateFormat6Choice `xml:"LpsdDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat6Choice `xml:"PmtDt,omitempty"`

	// Date/Time by which the account owner must instruct directly another party, for example to provide documentation to an issuer agent.
	ThirdPartyDeadline *DateFormat6Choice `xml:"ThrdPtyDdln,omitempty"`

	// Date/Time set by the issuer agent as a first early deadline by which the account owner must instruct directly another party, possibly giving the holder eligibility to incentives. For example, to provide documentation to an issuer agent.
	EarlyThirdPartyDeadline *DateFormat6Choice `xml:"EarlyThrdPtyDdln,omitempty"`

	// Date by which the depository stops monitoring activities of the event, for instance, accounting and tracking activities for due bills end.
	MarketClaimTrackingEndDate *DateFormat6Choice `xml:"MktClmTrckgEndDt,omitempty"`
}

func (c *CorporateActionDate1) AddAnnouncementDate() *DateFormat6Choice {
	c.AnnouncementDate = new(DateFormat6Choice)
	return c.AnnouncementDate
}

func (c *CorporateActionDate1) AddCertificationDeadline() *DateFormat6Choice {
	c.CertificationDeadline = new(DateFormat6Choice)
	return c.CertificationDeadline
}

func (c *CorporateActionDate1) AddCourtApprovalDate() *DateFormat6Choice {
	c.CourtApprovalDate = new(DateFormat6Choice)
	return c.CourtApprovalDate
}

func (c *CorporateActionDate1) AddEarlyClosingDate() *DateFormat6Choice {
	c.EarlyClosingDate = new(DateFormat6Choice)
	return c.EarlyClosingDate
}

func (c *CorporateActionDate1) AddEffectiveDate() *DateFormat6Choice {
	c.EffectiveDate = new(DateFormat6Choice)
	return c.EffectiveDate
}

func (c *CorporateActionDate1) AddEqualisationDate() *DateFormat6Choice {
	c.EqualisationDate = new(DateFormat6Choice)
	return c.EqualisationDate
}

func (c *CorporateActionDate1) AddFurtherDetailedAnnouncementDate() *DateFormat6Choice {
	c.FurtherDetailedAnnouncementDate = new(DateFormat6Choice)
	return c.FurtherDetailedAnnouncementDate
}

func (c *CorporateActionDate1) AddIndexFixingDate() *DateFormat6Choice {
	c.IndexFixingDate = new(DateFormat6Choice)
	return c.IndexFixingDate
}

func (c *CorporateActionDate1) AddLotteryDate() *DateFormat6Choice {
	c.LotteryDate = new(DateFormat6Choice)
	return c.LotteryDate
}

func (c *CorporateActionDate1) AddMaturityDate() *DateFormat6Choice {
	c.MaturityDate = new(DateFormat6Choice)
	return c.MaturityDate
}

func (c *CorporateActionDate1) AddMeetingDate() *DateFormat6Choice {
	c.MeetingDate = new(DateFormat6Choice)
	return c.MeetingDate
}

func (c *CorporateActionDate1) AddMarginFixingDate() *DateFormat6Choice {
	c.MarginFixingDate = new(DateFormat6Choice)
	return c.MarginFixingDate
}

func (c *CorporateActionDate1) AddProrationDate() *DateFormat6Choice {
	c.ProrationDate = new(DateFormat6Choice)
	return c.ProrationDate
}

func (c *CorporateActionDate1) AddRecordDate() *DateFormat6Choice {
	c.RecordDate = new(DateFormat6Choice)
	return c.RecordDate
}

func (c *CorporateActionDate1) AddRegistrationDeadline() *DateFormat6Choice {
	c.RegistrationDeadline = new(DateFormat6Choice)
	return c.RegistrationDeadline
}

func (c *CorporateActionDate1) AddResultsPublicationDate() *DateFormat6Choice {
	c.ResultsPublicationDate = new(DateFormat6Choice)
	return c.ResultsPublicationDate
}

func (c *CorporateActionDate1) AddDeadlineToSplit() *DateFormat6Choice {
	c.DeadlineToSplit = new(DateFormat6Choice)
	return c.DeadlineToSplit
}

func (c *CorporateActionDate1) AddDeadlineForTaxBreakdownInstruction() *DateFormat6Choice {
	c.DeadlineForTaxBreakdownInstruction = new(DateFormat6Choice)
	return c.DeadlineForTaxBreakdownInstruction
}

func (c *CorporateActionDate1) AddTradingSuspendedDate() *DateFormat6Choice {
	c.TradingSuspendedDate = new(DateFormat6Choice)
	return c.TradingSuspendedDate
}

func (c *CorporateActionDate1) AddUnconditionalDate() *DateFormat6Choice {
	c.UnconditionalDate = new(DateFormat6Choice)
	return c.UnconditionalDate
}

func (c *CorporateActionDate1) AddWhollyUnconditionalDate() *DateFormat6Choice {
	c.WhollyUnconditionalDate = new(DateFormat6Choice)
	return c.WhollyUnconditionalDate
}

func (c *CorporateActionDate1) AddExDividendDate() *DateFormat6Choice {
	c.ExDividendDate = new(DateFormat6Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate1) AddOfficialAnnouncementPublicationDate() *DateFormat6Choice {
	c.OfficialAnnouncementPublicationDate = new(DateFormat6Choice)
	return c.OfficialAnnouncementPublicationDate
}

func (c *CorporateActionDate1) AddSpecialExDate() *DateFormat6Choice {
	c.SpecialExDate = new(DateFormat6Choice)
	return c.SpecialExDate
}

func (c *CorporateActionDate1) AddGuaranteedParticipationDate() *DateFormat6Choice {
	c.GuaranteedParticipationDate = new(DateFormat6Choice)
	return c.GuaranteedParticipationDate
}

func (c *CorporateActionDate1) AddElectionToCounterpartyDeadline() *DateFormat6Choice {
	c.ElectionToCounterpartyDeadline = new(DateFormat6Choice)
	return c.ElectionToCounterpartyDeadline
}

func (c *CorporateActionDate1) AddLapsedDate() *DateFormat6Choice {
	c.LapsedDate = new(DateFormat6Choice)
	return c.LapsedDate
}

func (c *CorporateActionDate1) AddPaymentDate() *DateFormat6Choice {
	c.PaymentDate = new(DateFormat6Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate1) AddThirdPartyDeadline() *DateFormat6Choice {
	c.ThirdPartyDeadline = new(DateFormat6Choice)
	return c.ThirdPartyDeadline
}

func (c *CorporateActionDate1) AddEarlyThirdPartyDeadline() *DateFormat6Choice {
	c.EarlyThirdPartyDeadline = new(DateFormat6Choice)
	return c.EarlyThirdPartyDeadline
}

func (c *CorporateActionDate1) AddMarketClaimTrackingEndDate() *DateFormat6Choice {
	c.MarketClaimTrackingEndDate = new(DateFormat6Choice)
	return c.MarketClaimTrackingEndDate
}
