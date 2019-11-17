package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400104 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Document"`
	Message *MeetingInstructionV04 `xml:"MtgInstr"`
}

func (d *Document00400104) AddMessage() *MeetingInstructionV04 {
	d.Message = new(MeetingInstructionV04)
	return d.Message
}

// Scope
// A party holding the right to vote sends the MeetingInstruction message to an intermediary, the issuer or its agent to request the receiving party to act upon one or several instructions.
// Usage
// The MeetingInstruction message is used to register for a shareholders meeting, request blocking or registration of securities. It is used to assign a proxy, to specify the names of meeting attendees and to relay vote instructions per resolution electronically.
// The MeetingInstruction message may only be sent for one security, though several safekeeping places may be specified.
// Once the message is sent, it cannot be modified. It must be cancelled by a MeetingInstructionCancellationRequest. Only after receipt of a confirmed cancelled status via the MeetingInstructionStatus message, a new MeetingInstruction message can be sent.
type MeetingInstructionV04 struct {

	// Identifies the meeting instruction message.
	Identification *model.MessageIdentification1 `xml:"Id"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference4 `xml:"MtgRef"`

	// Party notifying the instructions.
	InstructingParty *model.PartyIdentification9Choice `xml:"InstgPty"`

	// Identifies the security for which the meeting is organised.
	SecurityIdentification *model.SecurityIdentification11 `xml:"SctyId"`

	// Identifies the position of the instructing party and the action that they want to take.
	Instruction []*model.Instruction2 `xml:"Instr"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (m *MeetingInstructionV04) AddIdentification() *model.MessageIdentification1 {
	m.Identification = new(model.MessageIdentification1)
	return m.Identification
}

func (m *MeetingInstructionV04) AddMeetingReference() *model.MeetingReference4 {
	m.MeetingReference = new(model.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingInstructionV04) AddInstructingParty() *model.PartyIdentification9Choice {
	m.InstructingParty = new(model.PartyIdentification9Choice)
	return m.InstructingParty
}

func (m *MeetingInstructionV04) AddSecurityIdentification() *model.SecurityIdentification11 {
	m.SecurityIdentification = new(model.SecurityIdentification11)
	return m.SecurityIdentification
}

func (m *MeetingInstructionV04) AddInstruction() *model.Instruction2 {
	newValue := new(model.Instruction2)
	m.Instruction = append(m.Instruction, newValue)
	return newValue
}

func (m *MeetingInstructionV04) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	m.Extension = append(m.Extension, newValue)
	return newValue
}
