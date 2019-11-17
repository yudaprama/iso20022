package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500104 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.005.001.04 Document"`
	Message *MeetingInstructionCancellationRequestV04 `xml:"MtgInstrCxlReq"`
}

func (d *Document00500104) AddMessage() *MeetingInstructionCancellationRequestV04 {
	d.Message = new(MeetingInstructionCancellationRequestV04)
	return d.Message
}

// Scope
// The MeetingInstructionCancellationRequest message is sent by the same party that sent the MeetingInstruction message. It is sent to request the cancellation of all instructions included in the original the MeetingInstruction message.
// Usage
// This message must be answered by a MeetingInstructionStatus message. Some instructions in the previously sent MeetingInstruction message may have already been executed and cannot be cancelled. This information should appear clearly in the status message.
type MeetingInstructionCancellationRequestV04 struct {

	// Uniquely identifies the cancellation request.
	Identification *model.MessageIdentification1 `xml:"Id"`

	// Identifies the instruction to be cancelled.
	PreviousReference *model.MessageIdentification `xml:"PrvsRef"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference4 `xml:"MtgRef,omitempty"`

	// Party requesting the cancellation.
	RequestingParty *model.PartyIdentification9Choice `xml:"RqstngPty,omitempty"`

	// Identifies the security for which the meeting is organised.
	SecurityIdentification *model.SecurityIdentification11 `xml:"SctyId,omitempty"`

	// Identifies the account and instructed positions for which the instruction cancellation request applies.
	InstructedPosition []*model.SafekeepingAccount4 `xml:"InstdPos,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (m *MeetingInstructionCancellationRequestV04) AddIdentification() *model.MessageIdentification1 {
	m.Identification = new(model.MessageIdentification1)
	return m.Identification
}

func (m *MeetingInstructionCancellationRequestV04) AddPreviousReference() *model.MessageIdentification {
	m.PreviousReference = new(model.MessageIdentification)
	return m.PreviousReference
}

func (m *MeetingInstructionCancellationRequestV04) AddMeetingReference() *model.MeetingReference4 {
	m.MeetingReference = new(model.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingInstructionCancellationRequestV04) AddRequestingParty() *model.PartyIdentification9Choice {
	m.RequestingParty = new(model.PartyIdentification9Choice)
	return m.RequestingParty
}

func (m *MeetingInstructionCancellationRequestV04) AddSecurityIdentification() *model.SecurityIdentification11 {
	m.SecurityIdentification = new(model.SecurityIdentification11)
	return m.SecurityIdentification
}

func (m *MeetingInstructionCancellationRequestV04) AddInstructedPosition() *model.SafekeepingAccount4 {
	newValue := new(model.SafekeepingAccount4)
	m.InstructedPosition = append(m.InstructedPosition, newValue)
	return newValue
}

func (m *MeetingInstructionCancellationRequestV04) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	m.Extension = append(m.Extension, newValue)
	return newValue
}
