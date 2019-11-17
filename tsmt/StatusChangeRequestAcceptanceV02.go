package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02700102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.027.001.02 Document"`
	Message *StatusChangeRequestAcceptanceV02 `xml:"StsChngReqAccptnc"`
}

func (d *Document02700102) AddMessage() *StatusChangeRequestAcceptanceV02 {
	d.Message = new(StatusChangeRequestAcceptanceV02)
	return d.Message
}

// Scope
// The StatusChangeRequestAcceptance message is sent by the party requested to accept or reject the request of a change in the status of a transaction to the matching application.
// This message is used to inform about the acceptance of a request to change the status of a transaction.
// Usage
// The StatusChangeRequestAcceptance message can be sent by the party requested to accept or reject a request to change the status of a transaction to inform that it accepts the request.
// The message can be sent in response to a StatusChangeRequestNotification message.
// The rejection of a request to change the status of a transaction can be achieved by sending a StatusChangeRequestRejection message.
type StatusChangeRequestAcceptanceV02 struct {

	// Identifies the acceptance message.
	AcceptanceIdentification *model.MessageIdentification1 `xml:"AccptncId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *model.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Specifies the status accepted.
	AcceptedStatus *model.TransactionStatus3 `xml:"AccptdSts"`
}

func (s *StatusChangeRequestAcceptanceV02) AddAcceptanceIdentification() *model.MessageIdentification1 {
	s.AcceptanceIdentification = new(model.MessageIdentification1)
	return s.AcceptanceIdentification
}

func (s *StatusChangeRequestAcceptanceV02) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	s.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusChangeRequestAcceptanceV02) AddSubmitterTransactionReference() *model.SimpleIdentificationInformation {
	s.SubmitterTransactionReference = new(model.SimpleIdentificationInformation)
	return s.SubmitterTransactionReference
}

func (s *StatusChangeRequestAcceptanceV02) AddAcceptedStatus() *model.TransactionStatus3 {
	s.AcceptedStatus = new(model.TransactionStatus3)
	return s.AcceptedStatus
}
