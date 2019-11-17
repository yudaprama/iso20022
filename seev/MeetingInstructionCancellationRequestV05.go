package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500105 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.005.001.05 Document"`
	Message *MeetingInstructionCancellationRequestV05 `xml:"MtgInstrCxlReq"`
}

func (d *Document00500105) AddMessage() *MeetingInstructionCancellationRequestV05 {
	d.Message = new(MeetingInstructionCancellationRequestV05)
	return d.Message
}

// Scope
// The MeetingInstructionCancellationRequest message is sent by the same party that sent the MeetingInstruction message. It is sent to request the cancellation of all instructions included in the original the MeetingInstruction message.
// Usage
// This message must be answered by a MeetingInstructionStatus message. Some instructions in the previously sent MeetingInstruction message may have already been executed and cannot be cancelled. This information should appear clearly in the status message.
// This message definition is intended for use with the Business Application Header (head.001.001.01).
type MeetingInstructionCancellationRequestV05 struct {

	// Identifies the instruction to be cancelled.
	PreviousReference *model.MessageIdentification `xml:"PrvsRef"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference7 `xml:"MtgRef,omitempty"`

	// Identifies the security for which the meeting is organised.
	FinancialInstrumentIdentification *model.SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Identifies the account and instructed positions for which the instruction cancellation request applies.
	InstructedPosition []*model.SafekeepingAccount6 `xml:"InstdPos,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MeetingInstructionCancellationRequestV05) AddPreviousReference() *model.MessageIdentification {
	m.PreviousReference = new(model.MessageIdentification)
	return m.PreviousReference
}

func (m *MeetingInstructionCancellationRequestV05) AddMeetingReference() *model.MeetingReference7 {
	m.MeetingReference = new(model.MeetingReference7)
	return m.MeetingReference
}

func (m *MeetingInstructionCancellationRequestV05) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	m.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return m.FinancialInstrumentIdentification
}

func (m *MeetingInstructionCancellationRequestV05) AddInstructedPosition() *model.SafekeepingAccount6 {
	newValue := new(model.SafekeepingAccount6)
	m.InstructedPosition = append(m.InstructedPosition, newValue)
	return newValue
}

func (m *MeetingInstructionCancellationRequestV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
