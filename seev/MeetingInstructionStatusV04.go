package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600104 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.04 Document"`
	Message *MeetingInstructionStatusV04 `xml:"MtgInstrSts"`
}

func (d *Document00600104) AddMessage() *MeetingInstructionStatusV04 {
	d.Message = new(MeetingInstructionStatusV04)
	return d.Message
}

// Scope
// The Receiver of the MeetingInstruction or MeetingInstructionCancellationRequest sends the MeetingInstructionStatus message to the Sender of these messages.
// The message gives the status of a complete message or of one or more specific instructions within the message.
// Usage
// The MeetingInstructionStatus message is used for four purposes.
// First, it provides the status on the processing of a MeetingInstructionCancellationRequest message, for example, whether the request message is rejected or accepted.
// Second, it is used to provide a global processing or rejection status of a MeetingInstruction message.
// Third, it is used to provide a detailed processing or rejection status of a MeetingInstruction message, for example, for each instruction in the MeetingInstruction message the processing or rejection status is individually reported by using the InstructionIdentification element. This identification allows the receiver of the status message to link the status confirmation to its original instruction.
// The blocking of securities should be confirmed via an MT 508 (Intra-Position Advice).
// Fourth, it is used as a reminder to request voting instructions. This is done by indicating NONREF in the Identification element of the InstructionIdentification component and by using the status code NotReceived in the ProcessingStatus.
type MeetingInstructionStatusV04 struct {

	// Identifies the meeting instruction status message.
	Identification *model.MessageIdentification1 `xml:"Id"`

	// Type of instruction.
	InstructionType *model.InstructionType1Choice `xml:"InstrTp"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference4 `xml:"MtgRef"`

	// Party reporting the status.
	ReportingParty *model.PartyIdentification9Choice `xml:"RptgPty"`

	// Identifies the securities for which the meeting is organised.
	SecurityIdentification *model.SecurityIdentification11 `xml:"SctyId"`

	// Type of instruction status.
	InstructionTypeStatus *model.InstructionTypeStatus1Choice `xml:"InstrTpSts"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (m *MeetingInstructionStatusV04) AddIdentification() *model.MessageIdentification1 {
	m.Identification = new(model.MessageIdentification1)
	return m.Identification
}

func (m *MeetingInstructionStatusV04) AddInstructionType() *model.InstructionType1Choice {
	m.InstructionType = new(model.InstructionType1Choice)
	return m.InstructionType
}

func (m *MeetingInstructionStatusV04) AddMeetingReference() *model.MeetingReference4 {
	m.MeetingReference = new(model.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingInstructionStatusV04) AddReportingParty() *model.PartyIdentification9Choice {
	m.ReportingParty = new(model.PartyIdentification9Choice)
	return m.ReportingParty
}

func (m *MeetingInstructionStatusV04) AddSecurityIdentification() *model.SecurityIdentification11 {
	m.SecurityIdentification = new(model.SecurityIdentification11)
	return m.SecurityIdentification
}

func (m *MeetingInstructionStatusV04) AddInstructionTypeStatus() *model.InstructionTypeStatus1Choice {
	m.InstructionTypeStatus = new(model.InstructionTypeStatus1Choice)
	return m.InstructionTypeStatus
}

func (m *MeetingInstructionStatusV04) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	m.Extension = append(m.Extension, newValue)
	return newValue
}
