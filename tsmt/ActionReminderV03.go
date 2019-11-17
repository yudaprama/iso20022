package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02400103 struct {
	XMLName xml.Name           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.024.001.03 Document"`
	Message *ActionReminderV03 `xml:"ActnRmndr"`
}

func (d *Document02400103) AddMessage() *ActionReminderV03 {
	d.Message = new(ActionReminderV03)
	return d.Message
}

// Scope
// The ActionReminder message is sent by the matching application to a party involved in a transaction that it is expecting to take an action.
// This message is used to remind a party of an action that it is expected to take.
// Usage
// The ActionReminder message can be sent by the matching application to remind a party that it is waiting for
// - the submission of a transaction initiation message (BaselineReSubmission message),
// or
// - the acceptance or rejection of mis-matched data sets (MisMatchAcceptance or MisMatchRejection message),
// or
// - the acceptance or rejection of an amendment request (AmendmentAcceptance or AmendmentRejection message),
// or
// - the acceptance or rejection of a status change request (StatusChangeRequestAcceptance or StatusChangeRequestRejection message),
// or
// - the acceptance or rejection of a status extension request (StatusExtensionAcceptance or StatusExtensionRejection message).
// - or
// - the acceptance or rejection of a request to accept role and baseline (RoleAndBaselineAcceptance or RoleAndBaselineRejection message).
type ActionReminderV03 struct {

	// Identifies the reminder message.
	ReminderIdentification *model.MessageIdentification1 `xml:"RmndrId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *model.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*model.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Identifies the message for which an action is required.
	MessageRequiringAction *model.MessageIdentification1 `xml:"MsgReqrngActn"`

	// Next processing step required.
	PendingRequestForAction *model.PendingActivity2 `xml:"PdgReqForActn"`
}

func (a *ActionReminderV03) AddReminderIdentification() *model.MessageIdentification1 {
	a.ReminderIdentification = new(model.MessageIdentification1)
	return a.ReminderIdentification
}

func (a *ActionReminderV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	a.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return a.TransactionIdentification
}

func (a *ActionReminderV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	a.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return a.EstablishedBaselineIdentification
}

func (a *ActionReminderV03) AddTransactionStatus() *model.TransactionStatus4 {
	a.TransactionStatus = new(model.TransactionStatus4)
	return a.TransactionStatus
}

func (a *ActionReminderV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	a.UserTransactionReference = append(a.UserTransactionReference, newValue)
	return newValue
}

func (a *ActionReminderV03) AddMessageRequiringAction() *model.MessageIdentification1 {
	a.MessageRequiringAction = new(model.MessageIdentification1)
	return a.MessageRequiringAction
}

func (a *ActionReminderV03) AddPendingRequestForAction() *model.PendingActivity2 {
	a.PendingRequestForAction = new(model.PendingActivity2)
	return a.PendingRequestForAction
}
