package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800102 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.02 Document"`
	Message *MeetingResultDisseminationV02 `xml:"MtgRsltDssmntn"`
}

func (d *Document00800102) AddMessage() *MeetingResultDisseminationV02 {
	d.Message = new(MeetingResultDisseminationV02)
	return d.Message
}

// Scope
// An issuer, its agent or an intermediary sends the MeetingResultDissemination message to another intermediary, to a party holding the right to vote, to a registered security holder or to a beneficial holder to provide information on the voting results of a shareholders meeting.
// Usage
// The MeetingResultDissemination message is used to provide the vote results per resolution. It may also provide information on the level of participation.
// This message is also used to notify an update or amendment to a previously sent MeetingResultDissemination message.
type MeetingResultDisseminationV02 struct {

	// Identifies the meeting dissemination notification message.
	MeetingResultDisseminationIdentification *model.MessageIdentification1 `xml:"MtgRsltDssmntnId"`

	// Information specific to an amemdment.
	Amendment *model.AmendInformation2 `xml:"Amdmnt,omitempty"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference3 `xml:"MtgRef"`

	// Party reporting the meeting results.
	ReportingParty *model.PartyIdentification9Choice `xml:"RptgPty"`

	// Identifies the securities for which the meeting is organised.
	Security []*model.SecurityPosition5 `xml:"Scty"`

	// Results per resolution.
	VoteResult []*model.Vote5 `xml:"VoteRslt"`

	// Information about the participation to the voting process.
	Participation *model.Participation2 `xml:"Prtcptn,omitempty"`

	// Information on where additionnal information can be received.
	AdditionalInformation *model.CommunicationAddress4 `xml:"AddtlInf,omitempty"`
}

func (m *MeetingResultDisseminationV02) AddMeetingResultDisseminationIdentification() *model.MessageIdentification1 {
	m.MeetingResultDisseminationIdentification = new(model.MessageIdentification1)
	return m.MeetingResultDisseminationIdentification
}

func (m *MeetingResultDisseminationV02) AddAmendment() *model.AmendInformation2 {
	m.Amendment = new(model.AmendInformation2)
	return m.Amendment
}

func (m *MeetingResultDisseminationV02) AddMeetingReference() *model.MeetingReference3 {
	m.MeetingReference = new(model.MeetingReference3)
	return m.MeetingReference
}

func (m *MeetingResultDisseminationV02) AddReportingParty() *model.PartyIdentification9Choice {
	m.ReportingParty = new(model.PartyIdentification9Choice)
	return m.ReportingParty
}

func (m *MeetingResultDisseminationV02) AddSecurity() *model.SecurityPosition5 {
	newValue := new(model.SecurityPosition5)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingResultDisseminationV02) AddVoteResult() *model.Vote5 {
	newValue := new(model.Vote5)
	m.VoteResult = append(m.VoteResult, newValue)
	return newValue
}

func (m *MeetingResultDisseminationV02) AddParticipation() *model.Participation2 {
	m.Participation = new(model.Participation2)
	return m.Participation
}

func (m *MeetingResultDisseminationV02) AddAdditionalInformation() *model.CommunicationAddress4 {
	m.AdditionalInformation = new(model.CommunicationAddress4)
	return m.AdditionalInformation
}
