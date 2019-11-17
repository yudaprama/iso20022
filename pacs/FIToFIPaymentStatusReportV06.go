package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200106 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.06 Document"`
	Message *FIToFIPaymentStatusReportV06 `xml:"FIToFIPmtStsRpt"`
}

func (d *Document00200106) AddMessage() *FIToFIPaymentStatusReportV06 {
	d.Message = new(FIToFIPaymentStatusReportV06)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionPaymentStatusReport message is sent by an instructed agent to the previous party in the payment chain. It is used to inform this party about the positive or negative status of an instruction (either single or file). It is also used to report on a pending instruction.
// Usage
// The FIToFIPaymentStatusReport message is exchanged between agents to provide status information about instructions previously sent. Its usage will always be governed by a bilateral agreement between the agents.
// The FIToFIPaymentStatusReport message can be used to provide information about the status (e.g. rejection, acceptance) of a credit transfer instruction, a direct debit instruction, as well as other intra-agent instructions (for example FIToFIPaymentCancellationRequest).
// The FIToFIPaymentStatusReport message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentStatusReport message can be used in domestic and cross-border scenarios.
type FIToFIPaymentStatusReportV06 struct {

	// Set of characteristics shared by all individual transactions included in the status report message.
	GroupHeader *model.GroupHeader53 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus []*model.OriginalGroupHeader1 `xml:"OrgnlGrpInfAndSts,omitempty"`

	// Information concerning the original transactions, to which the status report message refers.
	TransactionInformationAndStatus []*model.PaymentTransaction52 `xml:"TxInfAndSts,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *FIToFIPaymentStatusReportV06) AddGroupHeader() *model.GroupHeader53 {
	f.GroupHeader = new(model.GroupHeader53)
	return f.GroupHeader
}

func (f *FIToFIPaymentStatusReportV06) AddOriginalGroupInformationAndStatus() *model.OriginalGroupHeader1 {
	newValue := new(model.OriginalGroupHeader1)
	f.OriginalGroupInformationAndStatus = append(f.OriginalGroupInformationAndStatus, newValue)
	return newValue
}

func (f *FIToFIPaymentStatusReportV06) AddTransactionInformationAndStatus() *model.PaymentTransaction52 {
	newValue := new(model.PaymentTransaction52)
	f.TransactionInformationAndStatus = append(f.TransactionInformationAndStatus, newValue)
	return newValue
}

func (f *FIToFIPaymentStatusReportV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
