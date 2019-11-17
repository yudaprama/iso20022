package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03800101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.038.001.01 Document"`
	Message *InvoiceTaxReportStatusAdviceV01 `xml:"InvcTaxRptStsAdvc"`
}

func (d *Document03800101) AddMessage() *InvoiceTaxReportStatusAdviceV01 {
	d.Message = new(InvoiceTaxReportStatusAdviceV01)
	return d.Message
}

// The InvoiceTaxReportStatusAdvice message is sent by the matching application to the party from which it received a message.
// This message is used to acknowledge the InvoiceTaxReport message.
type InvoiceTaxReportStatusAdviceV01 struct {

	// Provides the status on the InvoiceTaxReport.
	StatusReportHeader *model.InvoiceTaxStatusReportHeader1 `xml:"StsRptHdr"`

	// Provides the status on an individual transaction and the related reason if required.
	TransactionStatus []*model.InvoiceTaxReportTransactionStatus1 `xml:"TxSts,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific
	// block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InvoiceTaxReportStatusAdviceV01) AddStatusReportHeader() *model.InvoiceTaxStatusReportHeader1 {
	i.StatusReportHeader = new(model.InvoiceTaxStatusReportHeader1)
	return i.StatusReportHeader
}

func (i *InvoiceTaxReportStatusAdviceV01) AddTransactionStatus() *model.InvoiceTaxReportTransactionStatus1 {
	newValue := new(model.InvoiceTaxReportTransactionStatus1)
	i.TransactionStatus = append(i.TransactionStatus, newValue)
	return newValue
}

func (i *InvoiceTaxReportStatusAdviceV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
