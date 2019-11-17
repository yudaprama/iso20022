package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04500102 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Document"`
	Message *ForwardIntentToPayNotificationV02 `xml:"FwdInttToPayNtfctn"`
}

func (d *Document04500102) AddMessage() *ForwardIntentToPayNotificationV02 {
	d.Message = new(ForwardIntentToPayNotificationV02)
	return d.Message
}

// Scope
// The ForwardIntentToPayNotification message is forwarded by the matching application from one primary bank to the other primary bank in order to provide details about a future payment.
// This message contains details about an intention to pay a certain amount, on a certain date, in relation to one or several transactions known to the matching application.
// Usage
// The ForwardIntentToPayNotification message is a copy of the IntentToPayNotification message received by the matching application and forwarded to the other primary bank for information. No response is expected.
type ForwardIntentToPayNotificationV02 struct {

	// Identifies the notification message.
	NotificationIdentification *model.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *model.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for the financial institutions involved in this transaction.
	UserTransactionReference []*model.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// The financial institution of the buyer, uniquely identified by its BIC.
	BuyerBank *model.BICIdentification1 `xml:"BuyrBk"`

	// The financial institution of the seller, uniquely identified by its BIC.
	SellerBank *model.BICIdentification1 `xml:"SellrBk"`

	// Provides the details of the intention to pay.
	IntentToPay *model.IntentToPay2 `xml:"InttToPay"`

	// Next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (f *ForwardIntentToPayNotificationV02) AddNotificationIdentification() *model.MessageIdentification1 {
	f.NotificationIdentification = new(model.MessageIdentification1)
	return f.NotificationIdentification
}

func (f *ForwardIntentToPayNotificationV02) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	f.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return f.TransactionIdentification
}

func (f *ForwardIntentToPayNotificationV02) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	f.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return f.EstablishedBaselineIdentification
}

func (f *ForwardIntentToPayNotificationV02) AddTransactionStatus() *model.TransactionStatus4 {
	f.TransactionStatus = new(model.TransactionStatus4)
	return f.TransactionStatus
}

func (f *ForwardIntentToPayNotificationV02) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	f.UserTransactionReference = append(f.UserTransactionReference, newValue)
	return newValue
}

func (f *ForwardIntentToPayNotificationV02) AddBuyerBank() *model.BICIdentification1 {
	f.BuyerBank = new(model.BICIdentification1)
	return f.BuyerBank
}

func (f *ForwardIntentToPayNotificationV02) AddSellerBank() *model.BICIdentification1 {
	f.SellerBank = new(model.BICIdentification1)
	return f.SellerBank
}

func (f *ForwardIntentToPayNotificationV02) AddIntentToPay() *model.IntentToPay2 {
	f.IntentToPay = new(model.IntentToPay2)
	return f.IntentToPay
}

func (f *ForwardIntentToPayNotificationV02) AddRequestForAction() *model.PendingActivity2 {
	f.RequestForAction = new(model.PendingActivity2)
	return f.RequestForAction
}
