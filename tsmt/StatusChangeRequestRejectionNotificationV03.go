package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03000103 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.030.001.03 Document"`
	Message *StatusChangeRequestRejectionNotificationV03 `xml:"StsChngReqRjctnNtfctn"`
}

func (d *Document03000103) AddMessage() *StatusChangeRequestRejectionNotificationV03 {
	d.Message = new(StatusChangeRequestRejectionNotificationV03)
	return d.Message
}

// Scope
// The StatusChangeRequestRejectionNotification message is sent by the matching application to the submitter of a request to change the status of a transaction.
// This message is used to inform about the rejection of a request to change the status of a transaction.
// Usage
// The StatusChangeRequestRejectionNotification message can be sent by the matching application to pass on information about the rejection of a request to change the status of a transaction that it has obtained through the receipt of a StatusChangeRequestRejection message.
// In order to pass on information about the acceptance of a request to change the status of a transaction the matching application sends a StatusChangeNotification message.
type StatusChangeRequestRejectionNotificationV03 struct {

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

	// Specifies the status that was rejected by the other party.
	RejectedStatusChange *model.TransactionStatus3 `xml:"RjctdStsChng"`

	// Reason why the user cannot accept the request.
	RejectionReason *model.Reason2 `xml:"RjctnRsn"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (s *StatusChangeRequestRejectionNotificationV03) AddNotificationIdentification() *model.MessageIdentification1 {
	s.NotificationIdentification = new(model.MessageIdentification1)
	return s.NotificationIdentification
}

func (s *StatusChangeRequestRejectionNotificationV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	s.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusChangeRequestRejectionNotificationV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	s.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return s.EstablishedBaselineIdentification
}

func (s *StatusChangeRequestRejectionNotificationV03) AddTransactionStatus() *model.TransactionStatus4 {
	s.TransactionStatus = new(model.TransactionStatus4)
	return s.TransactionStatus
}

func (s *StatusChangeRequestRejectionNotificationV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	s.UserTransactionReference = append(s.UserTransactionReference, newValue)
	return newValue
}

func (s *StatusChangeRequestRejectionNotificationV03) AddRejectedStatusChange() *model.TransactionStatus3 {
	s.RejectedStatusChange = new(model.TransactionStatus3)
	return s.RejectedStatusChange
}

func (s *StatusChangeRequestRejectionNotificationV03) AddRejectionReason() *model.Reason2 {
	s.RejectionReason = new(model.Reason2)
	return s.RejectionReason
}

func (s *StatusChangeRequestRejectionNotificationV03) AddRequestForAction() *model.PendingActivity2 {
	s.RequestForAction = new(model.PendingActivity2)
	return s.RequestForAction
}
