package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400105 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.05 Document"`
	Message *MeetingInstructionV05 `xml:"MtgInstr"`
}

func (d *Document00400105) AddMessage() *MeetingInstructionV05 {
	d.Message = new(MeetingInstructionV05)
	return d.Message
}

// Scope
// A party holding the right to vote sends the MeetingInstruction message to an intermediary, the issuer or its agent to request the receiving party to act upon one or several instructions.
// Usage
// The MeetingInstruction message is used to register for a shareholders meeting, request blocking or registration of securities. It is used to assign a proxy, to specify the names of meeting attendees and to relay vote instructions per resolution electronically.
// The MeetingInstruction message may only be sent for one security, though several safekeeping places may be specified.
// Once the message is sent, it cannot be modified. It must be cancelled by a MeetingInstructionCancellationRequest. Only after receipt of a confirmed cancelled status via the MeetingInstructionStatus message, a new MeetingInstruction message can be sent.
// This message definition is intended for use with the Business Application Header (head.001.001.01).
type MeetingInstructionV05 struct {

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference7 `xml:"MtgRef"`

	// Identifies the security for which the meeting is organised.
	FinancialInstrumentIdentification *model.SecurityIdentification14 `xml:"FinInstrmId"`

	// Identifies the position of the instructing party and the action that they want to take.
	Instruction []*model.Instruction3 `xml:"Instr"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MeetingInstructionV05) AddMeetingReference() *model.MeetingReference7 {
	m.MeetingReference = new(model.MeetingReference7)
	return m.MeetingReference
}

func (m *MeetingInstructionV05) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	m.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return m.FinancialInstrumentIdentification
}

func (m *MeetingInstructionV05) AddInstruction() *model.Instruction3 {
	newValue := new(model.Instruction3)
	m.Instruction = append(m.Instruction, newValue)
	return newValue
}

func (m *MeetingInstructionV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
