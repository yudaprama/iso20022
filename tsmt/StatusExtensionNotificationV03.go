package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03200103 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.032.001.03 Document"`
	Message *StatusExtensionNotificationV03 `xml:"StsXtnsnNtfctn"`
}

func (d *Document03200103) AddMessage() *StatusExtensionNotificationV03 {
	d.Message = new(StatusExtensionNotificationV03)
	return d.Message
}

// Scope
// The StatusExtensionNotification message is sent by the matching application to the parties involved in a request to extend the status of a transaction.
// This message is used to inform about the acceptance of a request to extend the status of a transaction.
// Usage
// The StatusExtensionNotification message can be sent by the matching application
// - to the submitter of a request to extend the status of a transaction to pass on information about the acceptance of the request that it has obtained through the receipt of an StatusExtensionAcceptance message.
// - to the accepter of a request to extend the status of a transaction to inform about the actual status of the transaction in response to a StatusExtensionAcceptance message.
// In order to pass on information about the rejection of a request to extend the status of a transaction the matching application sends a StatusExtensionRejectionNotification message.
type StatusExtensionNotificationV03 struct {

	// Identifies the notification message.
	NotificationIdentification *model.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*model.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Identifies the status that is being rolled over.
	ExtendedStatus *model.TransactionStatus5 `xml:"XtndedSts"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (s *StatusExtensionNotificationV03) AddNotificationIdentification() *model.MessageIdentification1 {
	s.NotificationIdentification = new(model.MessageIdentification1)
	return s.NotificationIdentification
}

func (s *StatusExtensionNotificationV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	s.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusExtensionNotificationV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	s.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return s.EstablishedBaselineIdentification
}

func (s *StatusExtensionNotificationV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	s.UserTransactionReference = append(s.UserTransactionReference, newValue)
	return newValue
}

func (s *StatusExtensionNotificationV03) AddExtendedStatus() *model.TransactionStatus5 {
	s.ExtendedStatus = new(model.TransactionStatus5)
	return s.ExtendedStatus
}

func (s *StatusExtensionNotificationV03) AddRequestForAction() *model.PendingActivity2 {
	s.RequestForAction = new(model.PendingActivity2)
	return s.RequestForAction
}
