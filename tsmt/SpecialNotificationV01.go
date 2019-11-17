package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04800101 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.048.001.01 Document"`
	Message *SpecialNotificationV01 `xml:"SpclNtfctn"`
}

func (d *Document04800101) AddMessage() *SpecialNotificationV01 {
	d.Message = new(SpecialNotificationV01)
	return d.Message
}

// Scope
// The SpecialNotification message is sent by the matching application to parties to the transaction, following the receipt of a SpecialRequest message.
// The SpecialRequest message is sent by a party to the transaction to the matching application if special circumstances are such that it cannot take part any longer to a specific transaction or that it cannot fulfill its role in the transaction.
// Usage
// The SpecialNotification message is sent to the parties to the transaction so that they can take appropriate action.
type SpecialNotificationV01 struct {

	// Identifies the notification message.
	NotificationIdentification *model.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *model.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for the financial institution that is the sender of the acknowledged message.
	UserTransactionReference []*model.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Party that has sent the special request.
	Initiator *model.BICIdentification1 `xml:"Initr"`

	// Notification received by the matching application and forwarded to another party.
	Notification *model.Notification1 `xml:"Ntfctn"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (s *SpecialNotificationV01) AddNotificationIdentification() *model.MessageIdentification1 {
	s.NotificationIdentification = new(model.MessageIdentification1)
	return s.NotificationIdentification
}

func (s *SpecialNotificationV01) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	s.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *SpecialNotificationV01) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	s.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return s.EstablishedBaselineIdentification
}

func (s *SpecialNotificationV01) AddTransactionStatus() *model.TransactionStatus4 {
	s.TransactionStatus = new(model.TransactionStatus4)
	return s.TransactionStatus
}

func (s *SpecialNotificationV01) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	s.UserTransactionReference = append(s.UserTransactionReference, newValue)
	return newValue
}

func (s *SpecialNotificationV01) AddInitiator() *model.BICIdentification1 {
	s.Initiator = new(model.BICIdentification1)
	return s.Initiator
}

func (s *SpecialNotificationV01) AddNotification() *model.Notification1 {
	s.Notification = new(model.Notification1)
	return s.Notification
}

func (s *SpecialNotificationV01) AddRequestForAction() *model.PendingActivity2 {
	s.RequestForAction = new(model.PendingActivity2)
	return s.RequestForAction
}
