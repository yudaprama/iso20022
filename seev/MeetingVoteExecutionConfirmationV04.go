package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700104 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.04 Document"`
	Message *MeetingVoteExecutionConfirmationV04 `xml:"MtgVoteExctnConf"`
}

func (d *Document00700104) AddMessage() *MeetingVoteExecutionConfirmationV04 {
	d.Message = new(MeetingVoteExecutionConfirmationV04)
	return d.Message
}

// Scope
// An issuer, its agent or an intermediary sends the MeetingVoteExecutionConfirmation message to confirm to the Sender of the MeetingInstruction message, the execution of their voting instruction.
// Usage
// This message is sent after the shareholders meeting has taken place. The Sender of this message confirms the execution of the vote at the meeting and confirms that the vote has been processed as instructed via the MeetingInstruction message.
// This messages is sent if the Sender of the MeetingInstruction message has requested such a confirmation or if market practice or regulation stipulates the need for a full audit trail.
type MeetingVoteExecutionConfirmationV04 struct {

	// Identifies the vote execution confirmation message.
	Identification *model.MessageIdentification1 `xml:"Id"`

	// Identifies the meeting instruction message.
	RelatedReference *model.MessageIdentification `xml:"RltdRef"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference4 `xml:"MtgRef"`

	// Party confirming the votes.
	ReportingParty *model.PartyIdentification9Choice `xml:"RptgPty"`

	// Identifies the securities for which the meeting is organised.
	SecurityIdentification *model.SecurityIdentification11 `xml:"SctyId"`

	// Specifies how a party has voted for each agenda item.
	VoteInstructions []*model.DetailedInstructionStatus9 `xml:"VoteInstrs"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (m *MeetingVoteExecutionConfirmationV04) AddIdentification() *model.MessageIdentification1 {
	m.Identification = new(model.MessageIdentification1)
	return m.Identification
}

func (m *MeetingVoteExecutionConfirmationV04) AddRelatedReference() *model.MessageIdentification {
	m.RelatedReference = new(model.MessageIdentification)
	return m.RelatedReference
}

func (m *MeetingVoteExecutionConfirmationV04) AddMeetingReference() *model.MeetingReference4 {
	m.MeetingReference = new(model.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingVoteExecutionConfirmationV04) AddReportingParty() *model.PartyIdentification9Choice {
	m.ReportingParty = new(model.PartyIdentification9Choice)
	return m.ReportingParty
}

func (m *MeetingVoteExecutionConfirmationV04) AddSecurityIdentification() *model.SecurityIdentification11 {
	m.SecurityIdentification = new(model.SecurityIdentification11)
	return m.SecurityIdentification
}

func (m *MeetingVoteExecutionConfirmationV04) AddVoteInstructions() *model.DetailedInstructionStatus9 {
	newValue := new(model.DetailedInstructionStatus9)
	m.VoteInstructions = append(m.VoteInstructions, newValue)
	return newValue
}

func (m *MeetingVoteExecutionConfirmationV04) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	m.Extension = append(m.Extension, newValue)
	return newValue
}
