package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03800101 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.038.001.01 Document"`
	Message *CaseStatusReportRequest `xml:"camt.038.001.01"`
}

func (d *Document03800101) AddMessage() *CaseStatusReportRequest {
	d.Message = new(CaseStatusReportRequest)
	return d.Message
}

// Scope
// The Case Status Report Request message is sent by a case creator or case assigner to a case assignee.
// This message is used to request the status of a case.
// Usage
// The Case Status Report Request message must be answered with a Case Status Report message. It can be used to request the status of a:
// - request to cancel payment case
// - request to modify payment case
// - unable to apply case
// - claim non receipt case
// The Case Status Report Request message covers one and only one case at a time. If a case creator or case assigner needs the status of several cases, then multiple Case Status Report Request messages must be sent.
// The Case Status Report Request message may be forwarded to subsequent case assignee(s) in the case processing chain.
// The processing of a case generates Notification Of Case Assignment and/or Resolution Of Investigation messages to the case creator/case assigner. They alone should provide collaborating parties sufficient information about the progress of the investigation. The Case Status Report Request must therefore only be used when no information has been received from the case assignee within the expected time frame.
// An agent may suspense an investigation by classifying it as overdue if he, after sending the request for status, does not receive any response after a long time. Agents may set their individual threshold wait-time.
type CaseStatusReportRequest struct {

	// Identifies the party requesting the status, the requested party, the identification and the date of the status.
	RequestHeader *model.ReportHeader `xml:"ReqHdr"`

	// Identifies the case.
	Case *model.Case `xml:"Case"`
}

func (c *CaseStatusReportRequest) AddRequestHeader() *model.ReportHeader {
	c.RequestHeader = new(model.ReportHeader)
	return c.RequestHeader
}

func (c *CaseStatusReportRequest) AddCase() *model.Case {
	c.Case = new(model.Case)
	return c.Case
}
