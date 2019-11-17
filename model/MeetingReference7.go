package model

// Identification of a meeting.
type MeetingReference7 struct {

	// Identification assigned to the general meeting by the party notifying the meeting. It must be unique for the party notifying the meeting.
	MeetingIdentification *Max35Text `xml:"MtgId,omitempty"`

	// Identification assigned to the meeting by the issuer. It must be unique for the issuer.
	IssuerMeetingIdentification *Max35Text `xml:"IssrMtgId,omitempty"`

	// Date and time at which the meeting will take place.
	MeetingDateAndTime *ISODateTime `xml:"MtgDtAndTm"`

	// Specifies the type of meeting for which  instructions are sent.
	Type *MeetingType3Code `xml:"Tp"`

	// Classifies the type of meeting.
	Classification *MeetingTypeClassification1Choice `xml:"Clssfctn,omitempty"`

	// Place of the company meeting for the scheduled meeting date.
	Location []*PostalAddress1 `xml:"Lctn,omitempty"`
}

func (m *MeetingReference7) SetMeetingIdentification(value string) {
	m.MeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingReference7) SetIssuerMeetingIdentification(value string) {
	m.IssuerMeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingReference7) SetMeetingDateAndTime(value string) {
	m.MeetingDateAndTime = (*ISODateTime)(&value)
}

func (m *MeetingReference7) SetType(value string) {
	m.Type = (*MeetingType3Code)(&value)
}

func (m *MeetingReference7) AddClassification() *MeetingTypeClassification1Choice {
	m.Classification = new(MeetingTypeClassification1Choice)
	return m.Classification
}

func (m *MeetingReference7) AddLocation() *PostalAddress1 {
	newValue := new(PostalAddress1)
	m.Location = append(m.Location, newValue)
	return newValue
}
