package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100101 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.001.001.01 Document"`
	Message *SecuritiesMessageRejection `xml:"semt.001.001.01"`
}

func (d *Document00100101) AddMessage() *SecuritiesMessageRejection {
	d.Message = new(SecuritiesMessageRejection)
	return d.Message
}

// Scope
// The SecuritiesMessageRejection message is sent by an executing party to the instructing party. Typically, this message is sent by an account servicer to the account owner.
// This message is used to reject a previously received message on which action cannot be taken.
// Usage
// The SecuritiesMessageRejection message can be sent for the following reasons:
// - the executing party does not recognise the linked reference, so the executing party cannot process the message
// - the instructing party should not have sent the message. This could be because the Receiver does not expect the message, eg, there is no SLA in place between the Sender and the Receiver.
// The SecuritiesMessageRejection message must not be used to reject an instruction message that cannot be processed for business reasons, eg, if information is missing in an instruction message or because securities are not available for settlement.
type SecuritiesMessageRejection struct {

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef"`

	// Reason to reject the message.
	Reason *model.RejectionReason1 `xml:"Rsn"`
}

func (s *SecuritiesMessageRejection) AddRelatedReference() *model.AdditionalReference2 {
	s.RelatedReference = new(model.AdditionalReference2)
	return s.RelatedReference
}

func (s *SecuritiesMessageRejection) AddReason() *model.RejectionReason1 {
	s.Reason = new(model.RejectionReason1)
	return s.Reason
}
