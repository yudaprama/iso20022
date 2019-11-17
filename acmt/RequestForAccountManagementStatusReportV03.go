package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500103 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.03 Document"`
	Message *RequestForAccountManagementStatusReportV03 `xml:"ReqForAcctMgmtStsRpt"`
}

func (d *Document00500103) AddMessage() *RequestForAccountManagementStatusReportV03 {
	d.Message = new(RequestForAccountManagementStatusReportV03)
	return d.Message
}

// Scope
// An account owner, for example, an investor or its designated agent, sends the RequestForAccountManagementStatusReport message to the account servicer, for example, a registrar, transfer agent or custodian bank to request the status of an AccountOpeningInstruction or an AccountModificationInstruction.
// Usage
// The RequestForAccountManagementStatusReport message is used to request the processing status of a previously sent AccountOpeningInstruction message or AccountModificationInstruction message for which a AccountDetailsConfirmation message has not yet been received.
type RequestForAccountManagementStatusReportV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Information to identify the account management instruction for which the status is requested.
	RequestDetails *model.AccountManagementMessageReference2 `xml:"ReqDtls"`
}

func (r *RequestForAccountManagementStatusReportV03) AddMessageIdentification() *model.MessageIdentification1 {
	r.MessageIdentification = new(model.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RequestForAccountManagementStatusReportV03) AddRequestDetails() *model.AccountManagementMessageReference2 {
	r.RequestDetails = new(model.AccountManagementMessageReference2)
	return r.RequestDetails
}
