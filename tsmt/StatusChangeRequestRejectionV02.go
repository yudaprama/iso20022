package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02900102 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.029.001.02 Document"`
	Message *StatusChangeRequestRejectionV02 `xml:"StsChngReqRjctn"`
}

func (d *Document02900102) AddMessage() *StatusChangeRequestRejectionV02 {
	d.Message = new(StatusChangeRequestRejectionV02)
	return d.Message
}

// Scope
// The StatusChangeRequestRejection message is sent by the party requested to accept or reject the request of a change in the status of a transaction to the matching application.
// This message is used to inform about the rejection of a request to change the status of a transaction.
// Usage
// The StatusChangeRequestRejection message can be sent by the party requested to accept or reject a request to change the status of a transaction to inform that it rejects the request.
// The message can be sent in response to a StatusChangeRequestNotification message.
// The acceptance of a request to change the status of a transaction can be achieved by sending a StatusChangeRequestAcceptance message.
type StatusChangeRequestRejectionV02 struct {

	// Identifies the rejection message.
	RejectionIdentification *model.MessageIdentification1 `xml:"RjctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *model.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Specifies the status rejected.
	RejectedStatusChange *model.TransactionStatus3 `xml:"RjctdStsChng"`

	// Reason why the user cannot accept the request.
	RejectionReason *model.Reason2 `xml:"RjctnRsn"`
}

func (s *StatusChangeRequestRejectionV02) AddRejectionIdentification() *model.MessageIdentification1 {
	s.RejectionIdentification = new(model.MessageIdentification1)
	return s.RejectionIdentification
}

func (s *StatusChangeRequestRejectionV02) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	s.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusChangeRequestRejectionV02) AddSubmitterTransactionReference() *model.SimpleIdentificationInformation {
	s.SubmitterTransactionReference = new(model.SimpleIdentificationInformation)
	return s.SubmitterTransactionReference
}

func (s *StatusChangeRequestRejectionV02) AddRejectedStatusChange() *model.TransactionStatus3 {
	s.RejectedStatusChange = new(model.TransactionStatus3)
	return s.RejectedStatusChange
}

func (s *StatusChangeRequestRejectionV02) AddRejectionReason() *model.Reason2 {
	s.RejectionReason = new(model.Reason2)
	return s.RejectionReason
}
