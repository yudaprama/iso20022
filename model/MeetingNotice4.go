package model

// Information about the shareholders meeting, specifying the participation requirements and the voting procedures. Alternatively, it may indicate where such information may be obtained.
type MeetingNotice4 struct {

	// Identification assigned to the general meeting by the party notifying the meeting. It must be unique for the party notifying the meeting.
	MeetingIdentification *Max35Text `xml:"MtgId,omitempty"`

	// Identification assigned to the meeting by the issuer. It must be unique for the issuer.
	IssuerMeetingIdentification *Max35Text `xml:"IssrMtgId,omitempty"`

	// Specifies the type of security holders meeting.
	Type *MeetingType3Code `xml:"Tp"`

	// Classifies the type of meeting.
	Classification *MeetingTypeClassification1Choice `xml:"Clssfctn,omitempty"`

	// Official meeting announcement date.
	AnnouncementDate *ISODate `xml:"AnncmntDt,omitempty"`

	// Indicates whether physical participation to the meeting is required in order to be allowed to vote.
	AttendanceRequired *YesNoIndicator `xml:"AttndncReqrd,omitempty"`

	// Indicates how to order the attendance card or to give notice of attendance.
	AttendanceConfirmationInformation *Max350Text `xml:"AttndncConfInf,omitempty"`

	// Date and time by which the beneficial owner or agent must notify of its intention to participate in the meeting. This deadline is set by an intermediary.
	AttendanceConfirmationDeadline *DateFormat29Choice `xml:"AttndncConfDdln,omitempty"`

	// Date and time by which the beneficial owner or agent must notify of its intention to participate in the meeting (STP mode). This deadline is set by an intermediary.
	AttendanceConfirmationSTPDeadline *DateFormat29Choice `xml:"AttndncConfSTPDdln,omitempty"`

	// Date and time by which the attendance to the meeting should be confirmed. This deadline is set by the issuer.
	AttendanceConfirmationMarketDeadline *DateFormat29Choice `xml:"AttndncConfMktDdln,omitempty"`

	// Address to use over the www (HTTP) service where additional information on the meeting may be found.
	AdditionalDocumentationURLAddress *Max256Text `xml:"AddtlDcmnttnURLAdr,omitempty"`

	// Additional procedural information about the general meeting, specifying the participation requirements and the voting procedures. Alternatively, this may indicate where such information may be obtained.
	AdditionalProcedureDetails []*AdditionalRights2 `xml:"AddtlPrcdrDtls,omitempty"`

	// Number of securities admitted to the vote, expressed as an amount and a currency.
	TotalNumberOfSecuritiesOutstanding *UnitOrFaceAmount1Choice `xml:"TtlNbOfSctiesOutsdng,omitempty"`

	// Number of rights admitted to the vote.
	TotalNumberOfVotingRights *Number `xml:"TtlNbOfVtngRghts,omitempty"`

	// Address where the information on the proxy should be sent.
	ProxyAppointmentNotificationAddress *PostalAddress1 `xml:"PrxyAppntmntNtfctnAdr,omitempty"`

	// Indicates whether a proxy is allowed.
	ProxyChoice *Proxy2Choice `xml:"PrxyChc,omitempty"`

	// Contact person at the party organising the meeting, at the issuer or at an intermediary.
	ContactPersonDetails []*MeetingContactPerson2 `xml:"CtctPrsnDtls,omitempty"`

	// Date on which the company publishes the results of its meeting.
	ResultPublicationDate *DateFormat3Choice `xml:"RsltPblctnDt,omitempty"`
}

func (m *MeetingNotice4) SetMeetingIdentification(value string) {
	m.MeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingNotice4) SetIssuerMeetingIdentification(value string) {
	m.IssuerMeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingNotice4) SetType(value string) {
	m.Type = (*MeetingType3Code)(&value)
}

func (m *MeetingNotice4) AddClassification() *MeetingTypeClassification1Choice {
	m.Classification = new(MeetingTypeClassification1Choice)
	return m.Classification
}

func (m *MeetingNotice4) SetAnnouncementDate(value string) {
	m.AnnouncementDate = (*ISODate)(&value)
}

func (m *MeetingNotice4) SetAttendanceRequired(value string) {
	m.AttendanceRequired = (*YesNoIndicator)(&value)
}

func (m *MeetingNotice4) SetAttendanceConfirmationInformation(value string) {
	m.AttendanceConfirmationInformation = (*Max350Text)(&value)
}

func (m *MeetingNotice4) AddAttendanceConfirmationDeadline() *DateFormat29Choice {
	m.AttendanceConfirmationDeadline = new(DateFormat29Choice)
	return m.AttendanceConfirmationDeadline
}

func (m *MeetingNotice4) AddAttendanceConfirmationSTPDeadline() *DateFormat29Choice {
	m.AttendanceConfirmationSTPDeadline = new(DateFormat29Choice)
	return m.AttendanceConfirmationSTPDeadline
}

func (m *MeetingNotice4) AddAttendanceConfirmationMarketDeadline() *DateFormat29Choice {
	m.AttendanceConfirmationMarketDeadline = new(DateFormat29Choice)
	return m.AttendanceConfirmationMarketDeadline
}

func (m *MeetingNotice4) SetAdditionalDocumentationURLAddress(value string) {
	m.AdditionalDocumentationURLAddress = (*Max256Text)(&value)
}

func (m *MeetingNotice4) AddAdditionalProcedureDetails() *AdditionalRights2 {
	newValue := new(AdditionalRights2)
	m.AdditionalProcedureDetails = append(m.AdditionalProcedureDetails, newValue)
	return newValue
}

func (m *MeetingNotice4) AddTotalNumberOfSecuritiesOutstanding() *UnitOrFaceAmount1Choice {
	m.TotalNumberOfSecuritiesOutstanding = new(UnitOrFaceAmount1Choice)
	return m.TotalNumberOfSecuritiesOutstanding
}

func (m *MeetingNotice4) SetTotalNumberOfVotingRights(value string) {
	m.TotalNumberOfVotingRights = (*Number)(&value)
}

func (m *MeetingNotice4) AddProxyAppointmentNotificationAddress() *PostalAddress1 {
	m.ProxyAppointmentNotificationAddress = new(PostalAddress1)
	return m.ProxyAppointmentNotificationAddress
}

func (m *MeetingNotice4) AddProxyChoice() *Proxy2Choice {
	m.ProxyChoice = new(Proxy2Choice)
	return m.ProxyChoice
}

func (m *MeetingNotice4) AddContactPersonDetails() *MeetingContactPerson2 {
	newValue := new(MeetingContactPerson2)
	m.ContactPersonDetails = append(m.ContactPersonDetails, newValue)
	return newValue
}

func (m *MeetingNotice4) AddResultPublicationDate() *DateFormat3Choice {
	m.ResultPublicationDate = new(DateFormat3Choice)
	return m.ResultPublicationDate
}
