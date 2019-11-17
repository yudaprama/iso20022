package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600103 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.006.001.03 Document"`
	Message *AmendmentAcceptanceNotificationV03 `xml:"AmdmntAccptncNtfctn"`
}

func (d *Document00600103) AddMessage() *AmendmentAcceptanceNotificationV03 {
	d.Message = new(AmendmentAcceptanceNotificationV03)
	return d.Message
}

// Scope
// The AmendmentAcceptanceNotification message is sent by the matching application to the requester of an amendment.
// This message is used to notify the acceptance of an amendment request.
// Usage
// The AmendmentAcceptanceNotification message can be sent by the matching application to pass on information about the acceptance of an amendment request that it has obtained through the receipt of an AmendmentAcceptance message.
// In order to pass on information about the rejection of an amendment request the matching application sends an AmendmentRejectionNotification message.
type AmendmentAcceptanceNotificationV03 struct {

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

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*model.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Reference to the identification of the delta report that contained the amendment.
	DeltaReportReference *model.MessageIdentification1 `xml:"DltaRptRef"`

	// Sequence number of the accepted baseline amendment.
	AcceptedAmendmentNumber *model.Count1 `xml:"AccptdAmdmntNb"`

	// Party that has accepted the amendment.
	Initiator *model.BICIdentification1 `xml:"Initr"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (a *AmendmentAcceptanceNotificationV03) AddNotificationIdentification() *model.MessageIdentification1 {
	a.NotificationIdentification = new(model.MessageIdentification1)
	return a.NotificationIdentification
}

func (a *AmendmentAcceptanceNotificationV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	a.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return a.TransactionIdentification
}

func (a *AmendmentAcceptanceNotificationV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	a.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return a.EstablishedBaselineIdentification
}

func (a *AmendmentAcceptanceNotificationV03) AddTransactionStatus() *model.TransactionStatus4 {
	a.TransactionStatus = new(model.TransactionStatus4)
	return a.TransactionStatus
}

func (a *AmendmentAcceptanceNotificationV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	a.UserTransactionReference = append(a.UserTransactionReference, newValue)
	return newValue
}

func (a *AmendmentAcceptanceNotificationV03) AddDeltaReportReference() *model.MessageIdentification1 {
	a.DeltaReportReference = new(model.MessageIdentification1)
	return a.DeltaReportReference
}

func (a *AmendmentAcceptanceNotificationV03) AddAcceptedAmendmentNumber() *model.Count1 {
	a.AcceptedAmendmentNumber = new(model.Count1)
	return a.AcceptedAmendmentNumber
}

func (a *AmendmentAcceptanceNotificationV03) AddInitiator() *model.BICIdentification1 {
	a.Initiator = new(model.BICIdentification1)
	return a.Initiator
}

func (a *AmendmentAcceptanceNotificationV03) AddRequestForAction() *model.PendingActivity2 {
	a.RequestForAction = new(model.PendingActivity2)
	return a.RequestForAction
}
