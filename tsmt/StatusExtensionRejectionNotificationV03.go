package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03400103 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.034.001.03 Document"`
	Message *StatusExtensionRejectionNotificationV03 `xml:"StsXtnsnRjctnNtfctn"`
}

func (d *Document03400103) AddMessage() *StatusExtensionRejectionNotificationV03 {
	d.Message = new(StatusExtensionRejectionNotificationV03)
	return d.Message
}

// Scope
// The StatusExtensionRejectionNotification message is sent by the matching application to the submitter of a request to extend the status of a transaction.
// This message is used to inform about the rejection of a request to extend the status of a transaction.
// Usage
// The StatusExtensionRejectionNotification message can be sent by the matching application to pass on information about the rejection of a request to extend the status of a transaction that it has obtained through the receipt of a StatusExtensionRejection message.
// In order to pass on information about the acceptance of a request to extend the status of a transaction the matching application sends a StatusExtensionNotification message
type StatusExtensionRejectionNotificationV03 struct {

	// Identifies the notification message.
	NotificationIdentification *model.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction that is not extended.
	NonExtendedStatus *model.TransactionStatus4 `xml:"NonXtndedSts"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*model.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Reason why the user cannot accept the request.
	RejectionReason *model.Reason2 `xml:"RjctnRsn"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (s *StatusExtensionRejectionNotificationV03) AddNotificationIdentification() *model.MessageIdentification1 {
	s.NotificationIdentification = new(model.MessageIdentification1)
	return s.NotificationIdentification
}

func (s *StatusExtensionRejectionNotificationV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	s.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusExtensionRejectionNotificationV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	s.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return s.EstablishedBaselineIdentification
}

func (s *StatusExtensionRejectionNotificationV03) AddNonExtendedStatus() *model.TransactionStatus4 {
	s.NonExtendedStatus = new(model.TransactionStatus4)
	return s.NonExtendedStatus
}

func (s *StatusExtensionRejectionNotificationV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	s.UserTransactionReference = append(s.UserTransactionReference, newValue)
	return newValue
}

func (s *StatusExtensionRejectionNotificationV03) AddRejectionReason() *model.Reason2 {
	s.RejectionReason = new(model.Reason2)
	return s.RejectionReason
}

func (s *StatusExtensionRejectionNotificationV03) AddRequestForAction() *model.PendingActivity2 {
	s.RequestForAction = new(model.PendingActivity2)
	return s.RequestForAction
}
