package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04100103 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.041.001.03 Document"`
	Message *TransactionReportV03 `xml:"TxRpt"`
}

func (d *Document04100103) AddMessage() *TransactionReportV03 {
	d.Message = new(TransactionReportV03)
	return d.Message
}

// Scope
// The TransactionReport message is sent by the matching application to the requester of a transaction report.
// This message is used to report on various details of transactions registered in the matching application.
// Usage
// The TransactionReport message can be sent by the matching application to report on various details of transactions that the requester of the report asked for. The message is sent in response to a TransactionReportRequest message.
type TransactionReportV03 struct {

	// Identifies the report.
	ReportIdentification *model.MessageIdentification1 `xml:"RptId"`

	// Reference to the previous message requesting the report.
	RelatedMessageReference *model.MessageIdentification1 `xml:"RltdMsgRef"`

	// Detailed description of the items that correspond to the parameters set in the request.
	ReportedItems []*model.TransactionReportItems3 `xml:"RptdItms,omitempty"`
}

func (t *TransactionReportV03) AddReportIdentification() *model.MessageIdentification1 {
	t.ReportIdentification = new(model.MessageIdentification1)
	return t.ReportIdentification
}

func (t *TransactionReportV03) AddRelatedMessageReference() *model.MessageIdentification1 {
	t.RelatedMessageReference = new(model.MessageIdentification1)
	return t.RelatedMessageReference
}

func (t *TransactionReportV03) AddReportedItems() *model.TransactionReportItems3 {
	newValue := new(model.TransactionReportItems3)
	t.ReportedItems = append(t.ReportedItems, newValue)
	return newValue
}
