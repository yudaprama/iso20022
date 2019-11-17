package model

// Information about the shareholders meeting, specifying the participation requirements and the voting procedures. Alternatively, it may indicate where such information may be obtained.
type MeetingNotice3 struct {

	// Identification assigned to a general meeting by the party notifying the meeting. It must be unique for the party notifying the meeting.
	MeetingIdentification *Max35Text `xml:"MtgId,omitempty"`

	// Identification assigned to a meeting by the issuer. It must be unique for the issuer.
	IssuerMeetingIdentification *Max35Text `xml:"IssrMtgId,omitempty"`

	// Specifies the type of security holders meeting.
	Type *MeetingType2Code `xml:"Tp"`

	// Classifies the type of meeting.
	Classification *MeetingTypeClassification1Choice `xml:"Clssfctn,omitempty"`

	// Official meeting announcement date.
	AnnouncementDate *ISODate `xml:"AnncmntDt,omitempty"`

	// Indicates whether physical participation to a meeting is required in order to be allowed to vote.
	AttendanceRequired *YesNoIndicator `xml:"AttndncReqrd"`

	// Indicates how to order the attendance card or to give notice of attendance.
	AttendanceConfirmationInformation *Max350Text `xml:"AttndncConfInf,omitempty"`

	// Date and time by which the beneficial owner or agent must notify of their intention to participate in a meeting. This deadline is set by an intermediary.
	AttendanceConfirmationDeadline *DateFormat2Choice `xml:"AttndncConfDdln,omitempty"`

	// Date and time by which the beneficial owner or agent must notify of their intention to participate in a meeting (STP mode). This deadline is set by an intermediary.
	AttendanceConfirmationSTPDeadline *DateFormat2Choice `xml:"AttndncConfSTPDdln,omitempty"`

	// Date and time by which the attendance to the meeting should be confirmed. This deadline is set by the issuer.
	AttendanceConfirmationMarketDeadline *DateFormat2Choice `xml:"AttndncConfMktDdln,omitempty"`

	// Address to use over the www (HTTP) service where addtional information on the meeting may be found.
	AdditionalDocumentationURLAddress *Max256Text `xml:"AddtlDcmnttnURLAdr,omitempty"`

	// Additional procedural information about the general meeting, specifying the participation requirements and the voting procedures. Alternatively, it may indicate where such information may be obtained.
	AdditionalProcedureDetails []*AdditionalRights1 `xml:"AddtlPrcdrDtls,omitempty"`

	// Number of securities admitted to the vote, expressed as an amount and a currency.
	TotalNumberOfSecuritiesOutstanding *ActiveCurrencyAndAmount `xml:"TtlNbOfSctiesOutsdng,omitempty"`

	// Number of rights admitted to the vote.
	TotalNumberOfVotingRights *Number `xml:"TtlNbOfVtngRghts,omitempty"`

	// Address where the information on the proxy should be sent.
	ProxyAppointmentNotificationAddress *PostalAddress1 `xml:"PrxyAppntmntNtfctnAdr,omitempty"`

	// Choice to signalize whether proxy is allowed.
	ProxyChoice *Proxy1Choice `xml:"PrxyChc,omitempty"`

	// Contact person at the party organising the meeting, at the issuer or at an intermediary.
	ContactPersonDetails []*MeetingContactPerson1 `xml:"CtctPrsnDtls,omitempty"`

	// Date on which a company publishes the results of its meeting.
	ResultPublicationDate *DateFormat3Choice `xml:"RsltPblctnDt,omitempty"`
}

func (m *MeetingNotice3) SetMeetingIdentification(value string) {
	m.MeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingNotice3) SetIssuerMeetingIdentification(value string) {
	m.IssuerMeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingNotice3) SetType(value string) {
	m.Type = (*MeetingType2Code)(&value)
}

func (m *MeetingNotice3) AddClassification() *MeetingTypeClassification1Choice {
	m.Classification = new(MeetingTypeClassification1Choice)
	return m.Classification
}

func (m *MeetingNotice3) SetAnnouncementDate(value string) {
	m.AnnouncementDate = (*ISODate)(&value)
}

func (m *MeetingNotice3) SetAttendanceRequired(value string) {
	m.AttendanceRequired = (*YesNoIndicator)(&value)
}

func (m *MeetingNotice3) SetAttendanceConfirmationInformation(value string) {
	m.AttendanceConfirmationInformation = (*Max350Text)(&value)
}

func (m *MeetingNotice3) AddAttendanceConfirmationDeadline() *DateFormat2Choice {
	m.AttendanceConfirmationDeadline = new(DateFormat2Choice)
	return m.AttendanceConfirmationDeadline
}

func (m *MeetingNotice3) AddAttendanceConfirmationSTPDeadline() *DateFormat2Choice {
	m.AttendanceConfirmationSTPDeadline = new(DateFormat2Choice)
	return m.AttendanceConfirmationSTPDeadline
}

func (m *MeetingNotice3) AddAttendanceConfirmationMarketDeadline() *DateFormat2Choice {
	m.AttendanceConfirmationMarketDeadline = new(DateFormat2Choice)
	return m.AttendanceConfirmationMarketDeadline
}

func (m *MeetingNotice3) SetAdditionalDocumentationURLAddress(value string) {
	m.AdditionalDocumentationURLAddress = (*Max256Text)(&value)
}

func (m *MeetingNotice3) AddAdditionalProcedureDetails() *AdditionalRights1 {
	newValue := new(AdditionalRights1)
	m.AdditionalProcedureDetails = append(m.AdditionalProcedureDetails, newValue)
	return newValue
}

func (m *MeetingNotice3) SetTotalNumberOfSecuritiesOutstanding(value, currency string) {
	m.TotalNumberOfSecuritiesOutstanding = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MeetingNotice3) SetTotalNumberOfVotingRights(value string) {
	m.TotalNumberOfVotingRights = (*Number)(&value)
}

func (m *MeetingNotice3) AddProxyAppointmentNotificationAddress() *PostalAddress1 {
	m.ProxyAppointmentNotificationAddress = new(PostalAddress1)
	return m.ProxyAppointmentNotificationAddress
}

func (m *MeetingNotice3) AddProxyChoice() *Proxy1Choice {
	m.ProxyChoice = new(Proxy1Choice)
	return m.ProxyChoice
}

func (m *MeetingNotice3) AddContactPersonDetails() *MeetingContactPerson1 {
	newValue := new(MeetingContactPerson1)
	m.ContactPersonDetails = append(m.ContactPersonDetails, newValue)
	return newValue
}

func (m *MeetingNotice3) AddResultPublicationDate() *DateFormat3Choice {
	m.ResultPublicationDate = new(DateFormat3Choice)
	return m.ResultPublicationDate
}
