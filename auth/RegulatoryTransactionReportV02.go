package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800102 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.008.001.02 Document"`
	Message *RegulatoryTransactionReportV02 `xml:"RgltryTxRpt"`
}

func (d *Document00800102) AddMessage() *RegulatoryTransactionReportV02 {
	d.Message = new(RegulatoryTransactionReportV02)
	return d.Message
}

// Scope
// A reporting institution, eg, an investment bank, sends the RegulatoryTransactionReport to a regulator or an intermediary (eg a reporting agent), to report the transaction details of a trade that has been executed on or off-exchange.
// Usage
// The message definition can be used to report more than one transaction. The message definition can also be used to specify, on a trade by trade basis, to which authorities the transaction report(s) need to be sent using the TransactionReportMarker.
type RegulatoryTransactionReportV02 struct {

	// Identification of the RegulatoryTransactionReport.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Provides details of the trade for which the transaction report is being sent.
	TransactionDetails []*model.TransactionDetails3 `xml:"TxDtls"`

	// Identification of the firm that is legally responsible for sending the transaction report.
	//
	ReportingInstitution *model.PartyIdentification23Choice `xml:"RptgInstn"`

	// Identifies the intermediary which is reporting on behalf on the ReportingInstitution. If there is a reporting chain, then the last party should override the previous one.
	ReportingAgent *model.PartyIdentification24Choice `xml:"RptgAgt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RegulatoryTransactionReportV02) AddIdentification() *model.DocumentIdentification8 {
	r.Identification = new(model.DocumentIdentification8)
	return r.Identification
}

func (r *RegulatoryTransactionReportV02) AddTransactionDetails() *model.TransactionDetails3 {
	newValue := new(model.TransactionDetails3)
	r.TransactionDetails = append(r.TransactionDetails, newValue)
	return newValue
}

func (r *RegulatoryTransactionReportV02) AddReportingInstitution() *model.PartyIdentification23Choice {
	r.ReportingInstitution = new(model.PartyIdentification23Choice)
	return r.ReportingInstitution
}

func (r *RegulatoryTransactionReportV02) AddReportingAgent() *model.PartyIdentification24Choice {
	r.ReportingAgent = new(model.PartyIdentification24Choice)
	return r.ReportingAgent
}

func (r *RegulatoryTransactionReportV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
