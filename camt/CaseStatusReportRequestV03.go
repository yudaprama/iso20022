package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03800103 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.038.001.03 Document"`
	Message *CaseStatusReportRequestV03 `xml:"CaseStsRptReq"`
}

func (d *Document03800103) AddMessage() *CaseStatusReportRequestV03 {
	d.Message = new(CaseStatusReportRequestV03)
	return d.Message
}

// Scope
// The CaseStatusReportRequest message is sent by a case creator or case assigner to a case assignee.
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
// An agent may suspend an investigation by classifying it as overdue if, this agent, after sending the request for the status of the investigation, does not receive any response after a long time. Agents may set their individual threshold wait-time.
type CaseStatusReportRequestV03 struct {

	// Identifies the party requesting the status, the requested party, the identification and the date of the status.
	RequestHeader *model.ReportHeader4 `xml:"ReqHdr"`

	// Identifies the investigation case.
	Case *model.Case3 `xml:"Case"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CaseStatusReportRequestV03) AddRequestHeader() *model.ReportHeader4 {
	c.RequestHeader = new(model.ReportHeader4)
	return c.RequestHeader
}

func (c *CaseStatusReportRequestV03) AddCase() *model.Case3 {
	c.Case = new(model.Case3)
	return c.Case
}

func (c *CaseStatusReportRequestV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
