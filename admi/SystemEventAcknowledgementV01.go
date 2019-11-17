package admi

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 Document"`
	Message *SystemEventAcknowledgementV01 `xml:"SysEvtAck"`
}

func (d *Document01100101) AddMessage() *SystemEventAcknowledgementV01 {
	d.Message = new(SystemEventAcknowledgementV01)
	return d.Message
}

// The SystemEventAcknowledgement message is sent by a participant of a central system to the central system to acknowledge the notification of an occurrence of an event in a central system.
//
type SystemEventAcknowledgementV01 struct {

	// Unique and unambiguous identifier for the message, as assigned by the sender.
	MessageIdentification *model.Max35Text `xml:"MsgId"`

	// Represents the original reference of the system event notification for which the acknowledgement is given, as assigned by the central system.
	OriginatorReference *model.Max35Text `xml:"OrgtrRef,omitempty"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *model.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Details of the system event being acknowledged.
	AcknowledgementDetails *model.Event1 `xml:"AckDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SystemEventAcknowledgementV01) SetMessageIdentification(value string) {
	s.MessageIdentification = (*model.Max35Text)(&value)
}

func (s *SystemEventAcknowledgementV01) SetOriginatorReference(value string) {
	s.OriginatorReference = (*model.Max35Text)(&value)
}

func (s *SystemEventAcknowledgementV01) SetSettlementSessionIdentifier(value string) {
	s.SettlementSessionIdentifier = (*model.Exact4AlphaNumericText)(&value)
}

func (s *SystemEventAcknowledgementV01) AddAcknowledgementDetails() *model.Event1 {
	s.AcknowledgementDetails = new(model.Event1)
	return s.AcknowledgementDetails
}

func (s *SystemEventAcknowledgementV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
