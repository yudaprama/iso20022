package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600103 struct {
	XMLName xml.Name        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.016.001.03 Document"`
	Message *ErrorReportV03 `xml:"ErrRpt"`
}

func (d *Document01600103) AddMessage() *ErrorReportV03 {
	d.Message = new(ErrorReportV03)
	return d.Message
}

// Scope
// The ErrorReport message is sent by the matching application to the party from which it received a message.
// This message is used to inform about the inability of the matching application to process a received message.
// Usage
// The ErrorReport message can be sent to a party from which the matching application received a message to inform about its inability to process the received message because
// - the syntax of the message is incorrect,or
// - the message content is inconsistent,or
// - according to the workflow implemented in the matching application, it did not expect the received message.
type ErrorReportV03 struct {

	// Identifies the report.
	ReportIdentification *model.MessageIdentification1 `xml:"RptId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId,omitempty"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *model.TransactionStatus4 `xml:"TxSts,omitempty"`

	// Reference to the transaction for the financial institution which is the sender of the rejected message.
	UserTransactionReference *model.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Reference to the identification of the rejected message.
	RejectedMessageReference *model.MessageIdentification1 `xml:"RjctdMsgRef,omitempty"`

	// Specifies the total number of errors identified in the rejected message.
	NumberOfErrors *model.Count1 `xml:"NbOfErrs"`

	// Describes the error that is the cause of the rejection.
	ErrorDescription []*model.ValidationResult3 `xml:"ErrDesc"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (e *ErrorReportV03) AddReportIdentification() *model.MessageIdentification1 {
	e.ReportIdentification = new(model.MessageIdentification1)
	return e.ReportIdentification
}

func (e *ErrorReportV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	e.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return e.TransactionIdentification
}

func (e *ErrorReportV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	e.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return e.EstablishedBaselineIdentification
}

func (e *ErrorReportV03) AddTransactionStatus() *model.TransactionStatus4 {
	e.TransactionStatus = new(model.TransactionStatus4)
	return e.TransactionStatus
}

func (e *ErrorReportV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	e.UserTransactionReference = new(model.DocumentIdentification5)
	return e.UserTransactionReference
}

func (e *ErrorReportV03) AddRejectedMessageReference() *model.MessageIdentification1 {
	e.RejectedMessageReference = new(model.MessageIdentification1)
	return e.RejectedMessageReference
}

func (e *ErrorReportV03) AddNumberOfErrors() *model.Count1 {
	e.NumberOfErrors = new(model.Count1)
	return e.NumberOfErrors
}

func (e *ErrorReportV03) AddErrorDescription() *model.ValidationResult3 {
	newValue := new(model.ValidationResult3)
	e.ErrorDescription = append(e.ErrorDescription, newValue)
	return newValue
}

func (e *ErrorReportV03) AddRequestForAction() *model.PendingActivity2 {
	e.RequestForAction = new(model.PendingActivity2)
	return e.RequestForAction
}
