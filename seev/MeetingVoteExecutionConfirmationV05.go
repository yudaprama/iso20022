package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700105 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.05 Document"`
	Message *MeetingVoteExecutionConfirmationV05 `xml:"MtgVoteExctnConf"`
}

func (d *Document00700105) AddMessage() *MeetingVoteExecutionConfirmationV05 {
	d.Message = new(MeetingVoteExecutionConfirmationV05)
	return d.Message
}

// Scope
// An issuer, its agent or an intermediary sends the MeetingVoteExecutionConfirmation message to confirm to the Sender of the MeetingInstruction message, the execution of their voting instruction.
// Usage
// This message is sent after the shareholders meeting has taken place. The Sender of this message confirms the execution of the vote at the meeting and confirms that the vote has been processed as instructed via the MeetingInstruction message.
// This messages is sent if the Sender of the MeetingInstruction message has requested such a confirmation or if market practice or regulation stipulates the need for a full audit trail.
// This message definition is intended for use with the Business Application Header (head.001.001.01).
type MeetingVoteExecutionConfirmationV05 struct {

	// Identifies the meeting instruction message.
	RelatedReference *model.MessageIdentification `xml:"RltdRef"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference7 `xml:"MtgRef"`

	// Identifies the securities for which the meeting is organised.
	FinancialInstrumentIdentification *model.SecurityIdentification14 `xml:"FinInstrmId"`

	// Specifies how a party has voted for each agenda item.
	VoteInstructions []*model.DetailedInstructionStatus10 `xml:"VoteInstrs"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MeetingVoteExecutionConfirmationV05) AddRelatedReference() *model.MessageIdentification {
	m.RelatedReference = new(model.MessageIdentification)
	return m.RelatedReference
}

func (m *MeetingVoteExecutionConfirmationV05) AddMeetingReference() *model.MeetingReference7 {
	m.MeetingReference = new(model.MeetingReference7)
	return m.MeetingReference
}

func (m *MeetingVoteExecutionConfirmationV05) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	m.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return m.FinancialInstrumentIdentification
}

func (m *MeetingVoteExecutionConfirmationV05) AddVoteInstructions() *model.DetailedInstructionStatus10 {
	newValue := new(model.DetailedInstructionStatus10)
	m.VoteInstructions = append(m.VoteInstructions, newValue)
	return newValue
}

func (m *MeetingVoteExecutionConfirmationV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
