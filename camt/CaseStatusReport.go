package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03900101 struct {
	XMLName xml.Name          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.039.001.01 Document"`
	Message *CaseStatusReport `xml:"camt.039.001.01"`
}

func (d *Document03900101) AddMessage() *CaseStatusReport {
	d.Message = new(CaseStatusReport)
	return d.Message
}

// Scope
// The Case Status Report message is sent by a case assignee to a case creator or case assigner.
// This message is used to report on the status of a case.
// Usage
// A Case Status Report message is sent in reply to a Case Status Report Request message. This message
// - covers one and only one case at a time. (If a case assignee needs to report on several cases, then multiple Case Status Report messages must be sent.)
// - may be forwarded to subsequent case assigner(s) until it reaches the end point
// - is able to indicate the fact that a case has been assigned to a party downstream in the payment processing chain
// - may not be used in place of a Resolution Of Investigation (except for the condition given in the next bullet point) or Notification Of Case Assignment message
// - may be skipped and replaced by a Resolution Of Investigation message if at the moment when the request for a investigation status arrives, the assignee has obtained a solution. (In this case a Resolution Of Investigation message can be sent in lieu of a Case Status Report and the case may be closed.)
type CaseStatusReport struct {

	// Specifies generic information about an investigation report.
	Header *model.ReportHeader `xml:"Hdr"`

	// Identifies the case.
	Case *model.Case `xml:"Case"`

	// Defines the status of the case.
	Status *model.CaseStatus `xml:"Sts"`

	// Identifies the last assignment performed.
	NewAssignment *model.CaseAssignment `xml:"NewAssgnmt,omitempty"`
}

func (c *CaseStatusReport) AddHeader() *model.ReportHeader {
	c.Header = new(model.ReportHeader)
	return c.Header
}

func (c *CaseStatusReport) AddCase() *model.Case {
	c.Case = new(model.Case)
	return c.Case
}

func (c *CaseStatusReport) AddStatus() *model.CaseStatus {
	c.Status = new(model.CaseStatus)
	return c.Status
}

func (c *CaseStatusReport) AddNewAssignment() *model.CaseAssignment {
	c.NewAssignment = new(model.CaseAssignment)
	return c.NewAssignment
}
