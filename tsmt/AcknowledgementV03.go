package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100103 struct {
	XMLName xml.Name            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.001.001.03 Document"`
	Message *AcknowledgementV03 `xml:"Ack"`
}

func (d *Document00100103) AddMessage() *AcknowledgementV03 {
	d.Message = new(AcknowledgementV03)
	return d.Message
}

// Scope
// The Acknowledgement message is sent by the matching application to the party from which it received a message.
// This message is used to acknowledge the receipt of a message by the matching application.
// Usage
// The Acknowledgement message can be sent to a party from which the matching application received a message to acknowledge the receipt of that message. The message is sent when the matching application does not send any other message in response to a received message.
type AcknowledgementV03 struct {

	// Identifies the acknowledgement message.
	AcknowledgementIdentification *model.MessageIdentification1 `xml:"AckId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId,omitempty"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *model.TransactionStatus4 `xml:"TxSts,omitempty"`

	// Reference to the transaction for the financial institution that is the sender of the acknowledged message.
	UserTransactionReference []*model.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Reference to the identification of the acknowledged message.
	AcknowledgedMessageReference *model.MessageIdentification1 `xml:"AckdMsgRef"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (a *AcknowledgementV03) AddAcknowledgementIdentification() *model.MessageIdentification1 {
	a.AcknowledgementIdentification = new(model.MessageIdentification1)
	return a.AcknowledgementIdentification
}

func (a *AcknowledgementV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	a.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return a.TransactionIdentification
}

func (a *AcknowledgementV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	a.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return a.EstablishedBaselineIdentification
}

func (a *AcknowledgementV03) AddTransactionStatus() *model.TransactionStatus4 {
	a.TransactionStatus = new(model.TransactionStatus4)
	return a.TransactionStatus
}

func (a *AcknowledgementV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	a.UserTransactionReference = append(a.UserTransactionReference, newValue)
	return newValue
}

func (a *AcknowledgementV03) AddAcknowledgedMessageReference() *model.MessageIdentification1 {
	a.AcknowledgedMessageReference = new(model.MessageIdentification1)
	return a.AcknowledgedMessageReference
}

func (a *AcknowledgementV03) AddRequestForAction() *model.PendingActivity2 {
	a.RequestForAction = new(model.PendingActivity2)
	return a.RequestForAction
}
