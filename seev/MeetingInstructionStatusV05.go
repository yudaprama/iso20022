package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600105 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Document"`
	Message *MeetingInstructionStatusV05 `xml:"MtgInstrSts"`
}

func (d *Document00600105) AddMessage() *MeetingInstructionStatusV05 {
	d.Message = new(MeetingInstructionStatusV05)
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
// This message definition is intended for use with the Business Application Header (head.001.001.01).
type MeetingInstructionStatusV05 struct {

	// Type of instruction.
	InstructionType *model.InstructionType1Choice `xml:"InstrTp"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference7 `xml:"MtgRef"`

	// Identifies the securities for which the meeting is organised.
	FinancialInstrumentIdentification *model.SecurityIdentification14 `xml:"FinInstrmId"`

	// Type of instruction status.
	InstructionTypeStatus *model.InstructionTypeStatus2Choice `xml:"InstrTpSts"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MeetingInstructionStatusV05) AddInstructionType() *model.InstructionType1Choice {
	m.InstructionType = new(model.InstructionType1Choice)
	return m.InstructionType
}

func (m *MeetingInstructionStatusV05) AddMeetingReference() *model.MeetingReference7 {
	m.MeetingReference = new(model.MeetingReference7)
	return m.MeetingReference
}

func (m *MeetingInstructionStatusV05) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	m.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return m.FinancialInstrumentIdentification
}

func (m *MeetingInstructionStatusV05) AddInstructionTypeStatus() *model.InstructionTypeStatus2Choice {
	m.InstructionTypeStatus = new(model.InstructionTypeStatus2Choice)
	return m.InstructionTypeStatus
}

func (m *MeetingInstructionStatusV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
