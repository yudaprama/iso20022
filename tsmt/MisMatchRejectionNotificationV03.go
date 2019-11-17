package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02300103 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.023.001.03 Document"`
	Message *MisMatchRejectionNotificationV03 `xml:"MisMtchRjctnNtfctn"`
}

func (d *Document02300103) AddMessage() *MisMatchRejectionNotificationV03 {
	d.Message = new(MisMatchRejectionNotificationV03)
	return d.Message
}

// Scope
// The MisMatchRejectionNotification message is sent by the matching application to the parties involved in the transaction.
// This message is used to notify the rejection of mis-matched data sets.
// Usage
// The MisMatchRejectionNotification message can be sent by the matching application to pass on the information about the rejection of mis-matched data sets that it has obtained through the receipt of an MisMatchRejection message.
// In order to pass on information about the acceptance of mis-matched data sets the matching application sends an MisMatchAcceptanceNotification message.
type MisMatchRejectionNotificationV03 struct {

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

	// Reference to the identification of the report that contained the difference.
	DataSetMatchReportReference *model.MessageIdentification1 `xml:"DataSetMtchRptRef"`

	// Specifies the reaons for rejecting the mismatch.
	RejectionReason *model.RejectionReason1Choice `xml:"RjctnRsn"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (m *MisMatchRejectionNotificationV03) AddNotificationIdentification() *model.MessageIdentification1 {
	m.NotificationIdentification = new(model.MessageIdentification1)
	return m.NotificationIdentification
}

func (m *MisMatchRejectionNotificationV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	m.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return m.TransactionIdentification
}

func (m *MisMatchRejectionNotificationV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	m.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return m.EstablishedBaselineIdentification
}

func (m *MisMatchRejectionNotificationV03) AddTransactionStatus() *model.TransactionStatus4 {
	m.TransactionStatus = new(model.TransactionStatus4)
	return m.TransactionStatus
}

func (m *MisMatchRejectionNotificationV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	m.UserTransactionReference = append(m.UserTransactionReference, newValue)
	return newValue
}

func (m *MisMatchRejectionNotificationV03) AddDataSetMatchReportReference() *model.MessageIdentification1 {
	m.DataSetMatchReportReference = new(model.MessageIdentification1)
	return m.DataSetMatchReportReference
}

func (m *MisMatchRejectionNotificationV03) AddRejectionReason() *model.RejectionReason1Choice {
	m.RejectionReason = new(model.RejectionReason1Choice)
	return m.RejectionReason
}

func (m *MisMatchRejectionNotificationV03) AddRequestForAction() *model.PendingActivity2 {
	m.RequestForAction = new(model.PendingActivity2)
	return m.RequestForAction
}
