package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500102 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.005.001.02 Document"`
	Message *MeetingInstructionCancellationRequestV02 `xml:"MtgInstrCxlReq"`
}

func (d *Document00500102) AddMessage() *MeetingInstructionCancellationRequestV02 {
	d.Message = new(MeetingInstructionCancellationRequestV02)
	return d.Message
}

// Scope
// The MeetingInstructionCancellationRequest message is sent by the same party that sent the MeetingInstruction message. It is sent to request the cancellation of all instructions included in the original the MeetingInstruction message.
// Usage
// This message must be answered by a MeetingInstructionStatus message. Some instructions in the previously sent MeetingInstruction message may have already been executed and cannot be cancelled. This information should appear clearly in the status message.
type MeetingInstructionCancellationRequestV02 struct {

	// Uniquely identifies the cancellation request.
	InstructionCancellationIdentification *model.MessageIdentification1 `xml:"InstrCxlId"`

	// Identifies the instruction to be cancelled.
	PreviousReference *model.MessageIdentification `xml:"PrvsRef"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference3 `xml:"MtgRef,omitempty"`

	// Party requesting the cancellation.
	RequestingParty *model.PartyIdentification9Choice `xml:"RqstngPty,omitempty"`

	// Identifies the security for which the meeting is organised.
	SecurityIdentification *model.SecurityIdentification3 `xml:"SctyId,omitempty"`

	// Identifies the account and instructed positions for which the instruction cancellation request applies.
	InstructedPosition []*model.SafekeepingAccount3 `xml:"InstdPos,omitempty"`
}

func (m *MeetingInstructionCancellationRequestV02) AddInstructionCancellationIdentification() *model.MessageIdentification1 {
	m.InstructionCancellationIdentification = new(model.MessageIdentification1)
	return m.InstructionCancellationIdentification
}

func (m *MeetingInstructionCancellationRequestV02) AddPreviousReference() *model.MessageIdentification {
	m.PreviousReference = new(model.MessageIdentification)
	return m.PreviousReference
}

func (m *MeetingInstructionCancellationRequestV02) AddMeetingReference() *model.MeetingReference3 {
	m.MeetingReference = new(model.MeetingReference3)
	return m.MeetingReference
}

func (m *MeetingInstructionCancellationRequestV02) AddRequestingParty() *model.PartyIdentification9Choice {
	m.RequestingParty = new(model.PartyIdentification9Choice)
	return m.RequestingParty
}

func (m *MeetingInstructionCancellationRequestV02) AddSecurityIdentification() *model.SecurityIdentification3 {
	m.SecurityIdentification = new(model.SecurityIdentification3)
	return m.SecurityIdentification
}

func (m *MeetingInstructionCancellationRequestV02) AddInstructedPosition() *model.SafekeepingAccount3 {
	newValue := new(model.SafekeepingAccount3)
	m.InstructedPosition = append(m.InstructedPosition, newValue)
	return newValue
}
