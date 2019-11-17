package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400102 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.02 Document"`
	Message *MeetingInstructionV02 `xml:"MtgInstr"`
}

func (d *Document00400102) AddMessage() *MeetingInstructionV02 {
	d.Message = new(MeetingInstructionV02)
	return d.Message
}

// Scope
// A party holding the right to vote sends the MeetingInstruction message to an intermediary, the issuer or its agent to request the receiving party to act upon one or several instructions.
// Usage
// The MeetingInstruction message is used to register for a shareholders meeting, request blocking or registration of securities. It is used to assign a proxy, to specify the names of meeting attendees and to relay vote instructions per resolution electronically.
// The MeetingInstruction message may only be sent for one security, though several safekeeping places may be specified.
// Once the message is sent, it cannot be modified. It must be cancelled by a MeetingInstructionCancellationRequest. Only after receipt of a confirmed cancelled status via the MeetingInstructionStatus message, a new MeetingInstruction message can be sent.
type MeetingInstructionV02 struct {

	// Identifies the meeting instruction message.
	MeetingInstructionIdentification *model.MessageIdentification1 `xml:"MtgInstrId"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference3 `xml:"MtgRef"`

	// Party notifying the instructions.
	InstructingParty *model.PartyIdentification9Choice `xml:"InstgPty"`

	// Identifies the security for which the meeting is organised.
	SecurityIdentification *model.SecurityIdentification3 `xml:"SctyId"`

	// Identifies the position of the instructing party and the action that they want to take.
	Instruction []*model.Instruction1 `xml:"Instr"`
}

func (m *MeetingInstructionV02) AddMeetingInstructionIdentification() *model.MessageIdentification1 {
	m.MeetingInstructionIdentification = new(model.MessageIdentification1)
	return m.MeetingInstructionIdentification
}

func (m *MeetingInstructionV02) AddMeetingReference() *model.MeetingReference3 {
	m.MeetingReference = new(model.MeetingReference3)
	return m.MeetingReference
}

func (m *MeetingInstructionV02) AddInstructingParty() *model.PartyIdentification9Choice {
	m.InstructingParty = new(model.PartyIdentification9Choice)
	return m.InstructingParty
}

func (m *MeetingInstructionV02) AddSecurityIdentification() *model.SecurityIdentification3 {
	m.SecurityIdentification = new(model.SecurityIdentification3)
	return m.SecurityIdentification
}

func (m *MeetingInstructionV02) AddInstruction() *model.Instruction1 {
	newValue := new(model.Instruction1)
	m.Instruction = append(m.Instruction, newValue)
	return newValue
}
