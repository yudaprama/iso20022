package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500103 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.005.001.03 Document"`
	Message *MeetingInstructionCancellationRequestV03 `xml:"MtgInstrCxlReq"`
}

func (d *Document00500103) AddMessage() *MeetingInstructionCancellationRequestV03 {
	d.Message = new(MeetingInstructionCancellationRequestV03)
	return d.Message
}

// Scope
// The MeetingInstructionCancellationRequest message is sent by the same party that sent the MeetingInstruction message. It is sent to request the cancellation of all instructions included in the original the MeetingInstruction message.
// Usage
// This message must be answered by a MeetingInstructionStatus message. Some instructions in the previously sent MeetingInstruction message may have already been executed and cannot be cancelled. This information should appear clearly in the status message.
type MeetingInstructionCancellationRequestV03 struct {

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
}

func (m *MeetingInstructionCancellationRequestV03) AddIdentification() *model.MessageIdentification1 {
	m.Identification = new(model.MessageIdentification1)
	return m.Identification
}

func (m *MeetingInstructionCancellationRequestV03) AddPreviousReference() *model.MessageIdentification {
	m.PreviousReference = new(model.MessageIdentification)
	return m.PreviousReference
}

func (m *MeetingInstructionCancellationRequestV03) AddMeetingReference() *model.MeetingReference4 {
	m.MeetingReference = new(model.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingInstructionCancellationRequestV03) AddRequestingParty() *model.PartyIdentification9Choice {
	m.RequestingParty = new(model.PartyIdentification9Choice)
	return m.RequestingParty
}

func (m *MeetingInstructionCancellationRequestV03) AddSecurityIdentification() *model.SecurityIdentification11 {
	m.SecurityIdentification = new(model.SecurityIdentification11)
	return m.SecurityIdentification
}

func (m *MeetingInstructionCancellationRequestV03) AddInstructedPosition() *model.SafekeepingAccount4 {
	newValue := new(model.SafekeepingAccount4)
	m.InstructedPosition = append(m.InstructedPosition, newValue)
	return newValue
}
