package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03100103 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.031.001.03 Document"`
	Message *StatusExtensionRequestAcceptanceV03 `xml:"StsXtnsnReqAccptnc"`
}

func (d *Document03100103) AddMessage() *StatusExtensionRequestAcceptanceV03 {
	d.Message = new(StatusExtensionRequestAcceptanceV03)
	return d.Message
}

// Scope
// The StatusExtensionRequestAcceptance message is sent by the party requested to accept or reject a request to extend the status of a transaction to the matching application.
// This message is used to inform about the acceptance of a request to extend the status of a transaction.
// Usage
// The StatusExtensionRequestAcceptance message can be sent by the party requested to accept or reject the request to extend the status of a transaction to inform that it accepts the request.
// The message can be sent in response to a StatusExtensionRequestNotification message.
// The rejection of a request to extend the status of a transaction can be achieved by sending a StatusExtensionRequestRejection message.
type StatusExtensionRequestAcceptanceV03 struct {

	// Identifies the acceptance message.
	AcceptanceIdentification *model.MessageIdentification1 `xml:"AccptncId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *model.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Identifies the status for which the user accepts an extension of the validity period..
	ExtendedStatus *model.TransactionStatus4 `xml:"XtndedSts"`
}

func (s *StatusExtensionRequestAcceptanceV03) AddAcceptanceIdentification() *model.MessageIdentification1 {
	s.AcceptanceIdentification = new(model.MessageIdentification1)
	return s.AcceptanceIdentification
}

func (s *StatusExtensionRequestAcceptanceV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	s.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusExtensionRequestAcceptanceV03) AddSubmitterTransactionReference() *model.SimpleIdentificationInformation {
	s.SubmitterTransactionReference = new(model.SimpleIdentificationInformation)
	return s.SubmitterTransactionReference
}

func (s *StatusExtensionRequestAcceptanceV03) AddExtendedStatus() *model.TransactionStatus4 {
	s.ExtendedStatus = new(model.TransactionStatus4)
	return s.ExtendedStatus
}
