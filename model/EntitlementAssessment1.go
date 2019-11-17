package model

// Specifies the parameters, such as dates, used to calculate the entitlement to vote at a general meeting.
type EntitlementAssessment1 struct {

	// Date by which the securities should be blocked. This deadline is set by an intermediary.
	SecuritiesBlockingDeadline *DateFormat2Choice `xml:"SctiesBlckgDdln,omitempty"`

	// Date by which the securities should be blocked. This deadline is set by the issuer. (STP mode)
	SecuritiesBlockingSTPDeadline *DateFormat2Choice `xml:"SctiesBlckgSTPDdln,omitempty"`

	// Date by which the securities should be blocked. This deadline is set by the issuer.
	SecuritiesBlockingMarketDeadline *DateFormat2Choice `xml:"SctiesBlckgMktDdln,omitempty"`

	// Date by which the blocking period for the securities should end.
	SecuritiesBlockingPeriodEndDate *ISODateTime `xml:"SctiesBlckgPrdEndDt,omitempty"`

	// Date at which the positions are struck to note which parties will receive the entitlement, e.g. record date, book close date...
	EntitlementFixingDate *DateFormat3Choice `xml:"EntitlmntFxgDt,omitempty"`

	// Date by which the securities have to be registered. This deadline is specified by an intermediary.
	RegistrationSecuritiesDeadline *DateFormat2Choice `xml:"RegnSctiesDdln,omitempty"`

	// Date by which the securities have to be registered. This deadline is specified by an intermediary (STP mode).
	RegistrationSecuritiesSTPDeadline *DateFormat2Choice `xml:"RegnSctiesSTPDdln,omitempty"`

	// Date by which the securities have to be registered. This deadline is set by the issuer.
	RegistrationSecuritiesMarketDeadline *DateFormat2Choice `xml:"RegnSctiesMktDdln,omitempty"`

	// Date by which the holder needs to register its intention to participate in the meeting process in order to be allowed to participate in the meeting event. This deadline is specified by an intermediary.
	RegistrationParticipationDeadline *DateFormat2Choice `xml:"RegnPrtcptnDdln,omitempty"`

	// Date by which the holder needs to register its intention to participate in the meeting process in order to be allowed to participate in the meeting event. This deadline is specified by an intermediary (STP mode).
	RegistrationParticipationSTPDeadline *DateFormat2Choice `xml:"RegnPrtcptnSTPDdln,omitempty"`

	// Date by which the holder needs to register its intention to participate in the meeting process in order to be allowed to participate in the meeting event. This deadline is set by the issuer.
	RegistrationParticipationMarketDeadline *DateFormat2Choice `xml:"RegnPrtcptnMktDdln,omitempty"`

	// Specifies the calculation method of the number of votes assigned to one security. This element should be used when the entitlement calculation rule is complex.
	EntitlementDescription *Max350Text `xml:"EntitlmntDesc,omitempty"`

	// Number of votes assigned to one security.
	EntitlementRatio *DecimalNumber `xml:"EntitlmntRatio,omitempty"`
}

func (e *EntitlementAssessment1) AddSecuritiesBlockingDeadline() *DateFormat2Choice {
	e.SecuritiesBlockingDeadline = new(DateFormat2Choice)
	return e.SecuritiesBlockingDeadline
}

func (e *EntitlementAssessment1) AddSecuritiesBlockingSTPDeadline() *DateFormat2Choice {
	e.SecuritiesBlockingSTPDeadline = new(DateFormat2Choice)
	return e.SecuritiesBlockingSTPDeadline
}

func (e *EntitlementAssessment1) AddSecuritiesBlockingMarketDeadline() *DateFormat2Choice {
	e.SecuritiesBlockingMarketDeadline = new(DateFormat2Choice)
	return e.SecuritiesBlockingMarketDeadline
}

func (e *EntitlementAssessment1) SetSecuritiesBlockingPeriodEndDate(value string) {
	e.SecuritiesBlockingPeriodEndDate = (*ISODateTime)(&value)
}

func (e *EntitlementAssessment1) AddEntitlementFixingDate() *DateFormat3Choice {
	e.EntitlementFixingDate = new(DateFormat3Choice)
	return e.EntitlementFixingDate
}

func (e *EntitlementAssessment1) AddRegistrationSecuritiesDeadline() *DateFormat2Choice {
	e.RegistrationSecuritiesDeadline = new(DateFormat2Choice)
	return e.RegistrationSecuritiesDeadline
}

func (e *EntitlementAssessment1) AddRegistrationSecuritiesSTPDeadline() *DateFormat2Choice {
	e.RegistrationSecuritiesSTPDeadline = new(DateFormat2Choice)
	return e.RegistrationSecuritiesSTPDeadline
}

func (e *EntitlementAssessment1) AddRegistrationSecuritiesMarketDeadline() *DateFormat2Choice {
	e.RegistrationSecuritiesMarketDeadline = new(DateFormat2Choice)
	return e.RegistrationSecuritiesMarketDeadline
}

func (e *EntitlementAssessment1) AddRegistrationParticipationDeadline() *DateFormat2Choice {
	e.RegistrationParticipationDeadline = new(DateFormat2Choice)
	return e.RegistrationParticipationDeadline
}

func (e *EntitlementAssessment1) AddRegistrationParticipationSTPDeadline() *DateFormat2Choice {
	e.RegistrationParticipationSTPDeadline = new(DateFormat2Choice)
	return e.RegistrationParticipationSTPDeadline
}

func (e *EntitlementAssessment1) AddRegistrationParticipationMarketDeadline() *DateFormat2Choice {
	e.RegistrationParticipationMarketDeadline = new(DateFormat2Choice)
	return e.RegistrationParticipationMarketDeadline
}

func (e *EntitlementAssessment1) SetEntitlementDescription(value string) {
	e.EntitlementDescription = (*Max350Text)(&value)
}

func (e *EntitlementAssessment1) SetEntitlementRatio(value string) {
	e.EntitlementRatio = (*DecimalNumber)(&value)
}
