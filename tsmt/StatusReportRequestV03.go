package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03800103 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.038.001.03 Document"`
	Message *StatusReportRequestV03 `xml:"StsRptReq"`
}

func (d *Document03800103) AddMessage() *StatusReportRequestV03 {
	d.Message = new(StatusReportRequestV03)
	return d.Message
}

// Scope
// The StatusReportRequest message is sent by a party involved in a transaction to the matching application.
// This message is used to request a report on the status of transactions registered in the matching application.
// Usage
// The StatusReportRequest message can be sent by either party involved in a transaction to request a report on the status and sub-statuses of all transactions that the requester is involved in.
// The application will respond to the request by sending a StatusReport message.
type StatusReportRequestV03 struct {

	// Identifies the request message.
	RequestIdentification *model.MessageIdentification1 `xml:"ReqId"`

	// Specifies the entities of the submitter for which the transactions have to be reported.
	EntitiesToBeReported []*model.BICIdentification1 `xml:"NttiesToBeRptd,omitempty"`
}

func (s *StatusReportRequestV03) AddRequestIdentification() *model.MessageIdentification1 {
	s.RequestIdentification = new(model.MessageIdentification1)
	return s.RequestIdentification
}

func (s *StatusReportRequestV03) AddEntitiesToBeReported() *model.BICIdentification1 {
	newValue := new(model.BICIdentification1)
	s.EntitiesToBeReported = append(s.EntitiesToBeReported, newValue)
	return newValue
}
