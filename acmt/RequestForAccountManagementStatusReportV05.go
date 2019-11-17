package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500105 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.05 Document"`
	Message *RequestForAccountManagementStatusReportV05 `xml:"ReqForAcctMgmtStsRpt"`
}

func (d *Document00500105) AddMessage() *RequestForAccountManagementStatusReportV05 {
	d.Message = new(RequestForAccountManagementStatusReportV05)
	return d.Message
}

// Scope
// The RequestForAccountManagementStatusReport message is sent by an account owner, for example, an investor or its designated agent, to the account servicer, for example, a registrar, transfer agent, custodian bank or securities depository  to request the status of an AccountOpeningInstruction,  GetAccountDetails or an AccountModificationInstruction.
// Usage
// The RequestForAccountManagementStatusReport message is used to request the processing status of a previously sent AccountOpeningInstruction, GetAccountDetails or an AccountModificationInstruction message for which an AccountDetailsConfirmation message has not yet been received.
type RequestForAccountManagementStatusReportV05 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Identifies the account for which the status of the account management instruction is requested.
	RequestDetails *model.AccountManagementMessageReference4 `xml:"ReqDtls"`
}

func (r *RequestForAccountManagementStatusReportV05) AddMessageIdentification() *model.MessageIdentification1 {
	r.MessageIdentification = new(model.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RequestForAccountManagementStatusReportV05) AddRequestDetails() *model.AccountManagementMessageReference4 {
	r.RequestDetails = new(model.AccountManagementMessageReference4)
	return r.RequestDetails
}
