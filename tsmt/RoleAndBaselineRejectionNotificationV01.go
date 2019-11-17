package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05200101 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 Document"`
	Message *RoleAndBaselineRejectionNotificationV01 `xml:"RoleAndBaselnRjctnNtfctn"`
}

func (d *Document05200101) AddMessage() *RoleAndBaselineRejectionNotificationV01 {
	d.Message = new(RoleAndBaselineRejectionNotificationV01)
	return d.Message
}

// Scope
// The RoleAndBaselineRejectionNotification message is sent by the matching application to the primary banks to inform about role and baseline rejection by a secondary bank.
// Usage
// The RoleAndBaselineRejectionNotification message is used to inform that a secondary bank has rejected the role and baseline. No response is expected.
type RoleAndBaselineRejectionNotificationV01 struct {

	// Identifies the notification message.
	NotificationIdentification *model.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction that is not extended.
	TransactionStatus *model.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*model.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Party that has rejected.
	Initiator *model.BICIdentification1 `xml:"Initr"`

	// Reason why the user cannot accept the request.
	RejectionReason *model.Reason2 `xml:"RjctnRsn,omitempty"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (r *RoleAndBaselineRejectionNotificationV01) AddNotificationIdentification() *model.MessageIdentification1 {
	r.NotificationIdentification = new(model.MessageIdentification1)
	return r.NotificationIdentification
}

func (r *RoleAndBaselineRejectionNotificationV01) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	r.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return r.TransactionIdentification
}

func (r *RoleAndBaselineRejectionNotificationV01) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	r.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return r.EstablishedBaselineIdentification
}

func (r *RoleAndBaselineRejectionNotificationV01) AddTransactionStatus() *model.TransactionStatus4 {
	r.TransactionStatus = new(model.TransactionStatus4)
	return r.TransactionStatus
}

func (r *RoleAndBaselineRejectionNotificationV01) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	r.UserTransactionReference = append(r.UserTransactionReference, newValue)
	return newValue
}

func (r *RoleAndBaselineRejectionNotificationV01) AddInitiator() *model.BICIdentification1 {
	r.Initiator = new(model.BICIdentification1)
	return r.Initiator
}

func (r *RoleAndBaselineRejectionNotificationV01) AddRejectionReason() *model.Reason2 {
	r.RejectionReason = new(model.Reason2)
	return r.RejectionReason
}

func (r *RoleAndBaselineRejectionNotificationV01) AddRequestForAction() *model.PendingActivity2 {
	r.RequestForAction = new(model.PendingActivity2)
	return r.RequestForAction
}
