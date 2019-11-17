package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800104 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.04 Document"`
	Message *MeetingResultDisseminationV04 `xml:"MtgRsltDssmntn"`
}

func (d *Document00800104) AddMessage() *MeetingResultDisseminationV04 {
	d.Message = new(MeetingResultDisseminationV04)
	return d.Message
}

// Scope
// An issuer, its agent or an intermediary sends the MeetingResultDissemination message to another intermediary, to a party holding the right to vote, to a registered security holder or to a beneficial holder to provide information on the voting results of a shareholders meeting.
// Usage
// The MeetingResultDissemination message is used to provide the vote results per resolution. It may also provide information on the level of participation.
// This message is also used to notify an update or amendment to a previously sent MeetingResultDissemination message.
type MeetingResultDisseminationV04 struct {

	// Identifies the meeting dissemination notification message.
	Identification *model.MessageIdentification1 `xml:"Id"`

	// Information specific to an amemdment.
	Amendment *model.AmendInformation2 `xml:"Amdmnt,omitempty"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference4 `xml:"MtgRef"`

	// Party reporting the meeting results.
	ReportingParty *model.PartyIdentification9Choice `xml:"RptgPty"`

	// Identifies the securities for which the meeting is organised.
	Security []*model.SecurityPosition6 `xml:"Scty"`

	// Results per resolution.
	VoteResult []*model.Vote5 `xml:"VoteRslt"`

	// Information about the participation to the voting process.
	Participation *model.Participation3 `xml:"Prtcptn,omitempty"`

	// Information on where additionnal information can be received.
	AdditionalInformation *model.CommunicationAddress4 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (m *MeetingResultDisseminationV04) AddIdentification() *model.MessageIdentification1 {
	m.Identification = new(model.MessageIdentification1)
	return m.Identification
}

func (m *MeetingResultDisseminationV04) AddAmendment() *model.AmendInformation2 {
	m.Amendment = new(model.AmendInformation2)
	return m.Amendment
}

func (m *MeetingResultDisseminationV04) AddMeetingReference() *model.MeetingReference4 {
	m.MeetingReference = new(model.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingResultDisseminationV04) AddReportingParty() *model.PartyIdentification9Choice {
	m.ReportingParty = new(model.PartyIdentification9Choice)
	return m.ReportingParty
}

func (m *MeetingResultDisseminationV04) AddSecurity() *model.SecurityPosition6 {
	newValue := new(model.SecurityPosition6)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingResultDisseminationV04) AddVoteResult() *model.Vote5 {
	newValue := new(model.Vote5)
	m.VoteResult = append(m.VoteResult, newValue)
	return newValue
}

func (m *MeetingResultDisseminationV04) AddParticipation() *model.Participation3 {
	m.Participation = new(model.Participation3)
	return m.Participation
}

func (m *MeetingResultDisseminationV04) AddAdditionalInformation() *model.CommunicationAddress4 {
	m.AdditionalInformation = new(model.CommunicationAddress4)
	return m.AdditionalInformation
}

func (m *MeetingResultDisseminationV04) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	m.Extension = append(m.Extension, newValue)
	return newValue
}
