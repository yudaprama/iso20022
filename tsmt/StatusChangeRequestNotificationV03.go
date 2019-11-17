package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02800103 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.028.001.03 Document"`
	Message *StatusChangeRequestNotificationV03 `xml:"StsChngReqNtfctn"`
}

func (d *Document02800103) AddMessage() *StatusChangeRequestNotificationV03 {
	d.Message = new(StatusChangeRequestNotificationV03)
	return d.Message
}

// Scope
// The StatusChangeRequestNotification message is sent by the matching application to the party requested to accept or reject the request of a change in the status of a transaction.
// This message is used to notify the request of a change in the status of a transaction.
// Usage
// The StatusChangeRequestNotification message can be sent by the matching application to pass on the information about the request of a change in the status of a transaction that it has obtained through the receipt of a StatusChangeRequest message.
type StatusChangeRequestNotificationV03 struct {

	// Identifies the notification message.
	NotificationIdentification *model.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *model.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*model.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Specifies the status that is proposed by the other party.
	ProposedStatusChange *model.TransactionStatus3 `xml:"PropsdStsChng"`

	// Specifies the reason for the request to change status.
	RequestReason *model.Reason2 `xml:"ReqRsn,omitempty"`

	// Party that has requested the status change.
	Initiator *model.BICIdentification1 `xml:"Initr"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (s *StatusChangeRequestNotificationV03) AddNotificationIdentification() *model.MessageIdentification1 {
	s.NotificationIdentification = new(model.MessageIdentification1)
	return s.NotificationIdentification
}

func (s *StatusChangeRequestNotificationV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	s.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusChangeRequestNotificationV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	s.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return s.EstablishedBaselineIdentification
}

func (s *StatusChangeRequestNotificationV03) AddTransactionStatus() *model.TransactionStatus4 {
	s.TransactionStatus = new(model.TransactionStatus4)
	return s.TransactionStatus
}

func (s *StatusChangeRequestNotificationV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	s.UserTransactionReference = append(s.UserTransactionReference, newValue)
	return newValue
}

func (s *StatusChangeRequestNotificationV03) AddProposedStatusChange() *model.TransactionStatus3 {
	s.ProposedStatusChange = new(model.TransactionStatus3)
	return s.ProposedStatusChange
}

func (s *StatusChangeRequestNotificationV03) AddRequestReason() *model.Reason2 {
	s.RequestReason = new(model.Reason2)
	return s.RequestReason
}

func (s *StatusChangeRequestNotificationV03) AddInitiator() *model.BICIdentification1 {
	s.Initiator = new(model.BICIdentification1)
	return s.Initiator
}

func (s *StatusChangeRequestNotificationV03) AddRequestForAction() *model.PendingActivity2 {
	s.RequestForAction = new(model.PendingActivity2)
	return s.RequestForAction
}
