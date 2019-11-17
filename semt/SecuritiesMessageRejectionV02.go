package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100102 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.001.001.02 Document"`
	Message *SecuritiesMessageRejectionV02 `xml:"SctiesMsgRjctnV02"`
}

func (d *Document00100102) AddMessage() *SecuritiesMessageRejectionV02 {
	d.Message = new(SecuritiesMessageRejectionV02)
	return d.Message
}

// Scope
// An account servicer, for example, a registrar, transfer agent or custodian bank, sends the SecuritiesMessageRejection message to the account owner, for example, an investor or its authorised agent, to reject a previously received message on which action cannot be taken.
// The message may also be sent by an executing party, for example, transfer agent to the instructing party, for example, investment manager or its authorised representative to reject a previously received message on which action cannot be taken.
// Usage
// The SecuritiesMessageRejection message is used for the following reasons:
// - the executing party does not recognise the linked reference, so the executing party cannot process the message
// - the instructing party should not have sent the message.
// Reasons that a receiver does not expect a message include no SLA in place between the Sender and the Receiver.
// The SecuritiesMessageRejection message must not be used to reject an instruction message that cannot be processed for business reasons, for example, if information is missing in an instruction message or because securities are not available for settlement.
type SecuritiesMessageRejectionV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef"`

	// Reason to reject the message.
	Reason *model.RejectionReason3 `xml:"Rsn"`
}

func (s *SecuritiesMessageRejectionV02) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SecuritiesMessageRejectionV02) AddRelatedReference() *model.AdditionalReference3 {
	s.RelatedReference = new(model.AdditionalReference3)
	return s.RelatedReference
}

func (s *SecuritiesMessageRejectionV02) AddReason() *model.RejectionReason3 {
	s.Reason = new(model.RejectionReason3)
	return s.Reason
}
