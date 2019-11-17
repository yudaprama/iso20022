package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600102 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.02 Document"`
	Message *MeetingInstructionStatusV02 `xml:"MtgInstrSts"`
}

func (d *Document00600102) AddMessage() *MeetingInstructionStatusV02 {
	d.Message = new(MeetingInstructionStatusV02)
	return d.Message
}

// Scope
// The Receiver of the MeetingInstruction or MeetingInstructionCancellationRequest sends the MeetingInstructionStatus message to the Sender of these messages.
// The message gives the status of a complete message or of one or more specific instructions within the message.
// Usage
// The MeetingInstructionStatus message is used for four purposes.
// First, it provides the status on the processing of a MeetingInstructionCancellationRequest message, ie, whether the request message is rejected or accepted.
// Second, it is used to provide a global processing or rejection status of a MeetingInstruction message.
// Third, it is used to provide a detailed processing or rejection status of a MeetingInstruction message, ie, for each instruction in the MeetingInstruction message the processing or rejection status is individually reported by using the InstructionIdentification element. This identification allows the receiver of the status message to link the status confirmation to its original instruction.
// The blocking of securities should be confirmed via an MT 508 (Intra-Position Advice).
// Fourth, it is used as a reminder to request voting instructions. This is done by indicating NONREF in the Identification element of the InstructionIdentification component and by using the status code NotReceived in the ProcessingStatus.
type MeetingInstructionStatusV02 struct {

	// Identifies the meeting instruction status message.
	MeetingInstructionStatusIdentification *model.MessageIdentification1 `xml:"MtgInstrStsId"`

	// Identifies the meeting instruction message for which the status is provided.
	InstructionIdentification *model.MessageIdentification `xml:"InstrId"`

	// Identifies the meeting instruction cancellation request message for which the status is provided.
	InstructionCancellationIdentification *model.MessageIdentification `xml:"InstrCxlId"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference3 `xml:"MtgRef"`

	// Party reporting the status.
	ReportingParty *model.PartyIdentification9Choice `xml:"RptgPty"`

	// Identifies the securities for which the meeting is organised.
	SecurityIdentification *model.SecurityIdentification3 `xml:"SctyId"`

	// Status applying to the instruction request received. The instruction is identified by the InstructionIdentification.
	InstructionStatus *model.InstructionStatus1Choice `xml:"InstrSts"`

	// Status applying to the instruction cancellation request received. The instruction cancellation is identified by the InstructionCancellationIdentification.
	CancellationStatus *model.CancellationStatus1Choice `xml:"CxlSts"`
}

func (m *MeetingInstructionStatusV02) AddMeetingInstructionStatusIdentification() *model.MessageIdentification1 {
	m.MeetingInstructionStatusIdentification = new(model.MessageIdentification1)
	return m.MeetingInstructionStatusIdentification
}

func (m *MeetingInstructionStatusV02) AddInstructionIdentification() *model.MessageIdentification {
	m.InstructionIdentification = new(model.MessageIdentification)
	return m.InstructionIdentification
}

func (m *MeetingInstructionStatusV02) AddInstructionCancellationIdentification() *model.MessageIdentification {
	m.InstructionCancellationIdentification = new(model.MessageIdentification)
	return m.InstructionCancellationIdentification
}

func (m *MeetingInstructionStatusV02) AddMeetingReference() *model.MeetingReference3 {
	m.MeetingReference = new(model.MeetingReference3)
	return m.MeetingReference
}

func (m *MeetingInstructionStatusV02) AddReportingParty() *model.PartyIdentification9Choice {
	m.ReportingParty = new(model.PartyIdentification9Choice)
	return m.ReportingParty
}

func (m *MeetingInstructionStatusV02) AddSecurityIdentification() *model.SecurityIdentification3 {
	m.SecurityIdentification = new(model.SecurityIdentification3)
	return m.SecurityIdentification
}

func (m *MeetingInstructionStatusV02) AddInstructionStatus() *model.InstructionStatus1Choice {
	m.InstructionStatus = new(model.InstructionStatus1Choice)
	return m.InstructionStatus
}

func (m *MeetingInstructionStatusV02) AddCancellationStatus() *model.CancellationStatus1Choice {
	m.CancellationStatus = new(model.CancellationStatus1Choice)
	return m.CancellationStatus
}
