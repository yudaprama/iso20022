package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700102 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Document"`
	Message *MeetingVoteExecutionConfirmationV02 `xml:"MtgVoteExctnConf"`
}

func (d *Document00700102) AddMessage() *MeetingVoteExecutionConfirmationV02 {
	d.Message = new(MeetingVoteExecutionConfirmationV02)
	return d.Message
}

// Scope
// An issuer, its agent or an intermediary sends the MeetingVoteExecutionConfirmation message to confirm to the Sender of the MeetingInstruction message, the execution of their voting instruction.
// Usage
// This message is sent after the shareholders meeting has taken place. The Sender of this message confirms the execution of the vote at the meeting and confirms that the vote has been processed as instructed via the MeetingInstruction message.
// This messages is sent if the Sender of the MeetingInstruction message has requested such a confirmation or if market practice or regulation stipulates the need for a full audit trail.
type MeetingVoteExecutionConfirmationV02 struct {

	// Identifies the vote execution confirmation message.
	VoteExecutionConfirmationIdentification *model.MessageIdentification1 `xml:"VoteExctnConfId"`

	// Identifies the meeting instruction message.
	RelatedReference *model.MessageIdentification `xml:"RltdRef"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference3 `xml:"MtgRef"`

	// Party confirming the votes.
	ReportingParty *model.PartyIdentification9Choice `xml:"RptgPty"`

	// Identifies the securities for which the meeting is organised.
	SecurityIdentification *model.SecurityIdentification3 `xml:"SctyId"`

	// Specifies how a party has voted for each agenda item.
	VoteInstruction []*model.DetailedInstructionStatus2 `xml:"VoteInstr"`
}

func (m *MeetingVoteExecutionConfirmationV02) AddVoteExecutionConfirmationIdentification() *model.MessageIdentification1 {
	m.VoteExecutionConfirmationIdentification = new(model.MessageIdentification1)
	return m.VoteExecutionConfirmationIdentification
}

func (m *MeetingVoteExecutionConfirmationV02) AddRelatedReference() *model.MessageIdentification {
	m.RelatedReference = new(model.MessageIdentification)
	return m.RelatedReference
}

func (m *MeetingVoteExecutionConfirmationV02) AddMeetingReference() *model.MeetingReference3 {
	m.MeetingReference = new(model.MeetingReference3)
	return m.MeetingReference
}

func (m *MeetingVoteExecutionConfirmationV02) AddReportingParty() *model.PartyIdentification9Choice {
	m.ReportingParty = new(model.PartyIdentification9Choice)
	return m.ReportingParty
}

func (m *MeetingVoteExecutionConfirmationV02) AddSecurityIdentification() *model.SecurityIdentification3 {
	m.SecurityIdentification = new(model.SecurityIdentification3)
	return m.SecurityIdentification
}

func (m *MeetingVoteExecutionConfirmationV02) AddVoteInstruction() *model.DetailedInstructionStatus2 {
	newValue := new(model.DetailedInstructionStatus2)
	m.VoteInstruction = append(m.VoteInstruction, newValue)
	return newValue
}
