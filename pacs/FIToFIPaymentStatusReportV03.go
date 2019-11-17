package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200103 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.03 Document"`
	Message *FIToFIPaymentStatusReportV03 `xml:"FIToFIPmtStsRpt"`
}

func (d *Document00200103) AddMessage() *FIToFIPaymentStatusReportV03 {
	d.Message = new(FIToFIPaymentStatusReportV03)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionPaymentStatusReport message is sent by an instructed agent to the previous party in the payment chain. It is used to inform this party about the positive or negative status of an instruction (either single or file). It is also used to report on a pending instruction.
// Usage
// The FIToFIPaymentStatusReport message is exchanged between agents to provide status information about instructions previously sent. Its usage will always be governed by a bilateral agreement between the agents.
// The FIToFIPaymentStatusReport message can be used to provide information about the status (e.g. rejection, acceptance) of a credit transfer instruction, a direct debit instruction, as well as other intra-agent instructions (for example FIToFIPaymentCancellationRequest).
// The FIToFIPaymentStatusReport message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentStatusReport message can be used in domestic and cross-border scenarios.
type FIToFIPaymentStatusReportV03 struct {

	// Set of characteristics shared by all individual transactions included in the status report message.
	GroupHeader *model.GroupHeader37 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus *model.OriginalGroupInformation20 `xml:"OrgnlGrpInfAndSts"`

	// Information concerning the original transactions, to which the status report message refers.
	TransactionInformationAndStatus []*model.PaymentTransactionInformation26 `xml:"TxInfAndSts,omitempty"`
}

func (f *FIToFIPaymentStatusReportV03) AddGroupHeader() *model.GroupHeader37 {
	f.GroupHeader = new(model.GroupHeader37)
	return f.GroupHeader
}

func (f *FIToFIPaymentStatusReportV03) AddOriginalGroupInformationAndStatus() *model.OriginalGroupInformation20 {
	f.OriginalGroupInformationAndStatus = new(model.OriginalGroupInformation20)
	return f.OriginalGroupInformationAndStatus
}

func (f *FIToFIPaymentStatusReportV03) AddTransactionInformationAndStatus() *model.PaymentTransactionInformation26 {
	newValue := new(model.PaymentTransactionInformation26)
	f.TransactionInformationAndStatus = append(f.TransactionInformationAndStatus, newValue)
	return newValue
}
