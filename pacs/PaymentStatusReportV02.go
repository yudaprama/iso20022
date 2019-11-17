package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200102 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.02 Document"`
	Message *PaymentStatusReportV02 `xml:"pacs.002.001.02"`
}

func (d *Document00200102) AddMessage() *PaymentStatusReportV02 {
	d.Message = new(PaymentStatusReportV02)
	return d.Message
}

// Scope
// The PaymentStatusReport message is sent by an instructed agent to the previous party in the payment chain. It is used to inform this party about the positive or negative status of an instruction (either single or file). It is also used to report on a pending instruction.
// Usage
// The PaymentStatusReport message is exchanged between agents to provide status information about instructions previously sent. Its usage will always be governed by a bilateral agreement between the agents.
// The PaymentStatusReport message can be used to provide information about the status (e.g. rejection, acceptance) of a credit transfer instruction, a direct debit instruction, as well as other intra-agent instructions (e.g. PaymentCancellationRequest).
// The PaymentStatusReport message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The PaymentStatusReport message can be used in domestic and cross-border scenarios.
// The PaymentStatusReport message exchanged between agents is identified in the schema as follows:
// urn:iso:std:iso:20022:tech:xsd:pacs.002.001.02
type PaymentStatusReportV02 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader5 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the message refers to.
	OriginalGroupInformationAndStatus *model.OriginalGroupInformation1 `xml:"OrgnlGrpInfAndSts"`

	// Information concerning the original transactions, to which the status report message refers.
	TransactionInformationAndStatus []*model.PaymentTransactionInformation1 `xml:"TxInfAndSts,omitempty"`
}

func (p *PaymentStatusReportV02) AddGroupHeader() *model.GroupHeader5 {
	p.GroupHeader = new(model.GroupHeader5)
	return p.GroupHeader
}

func (p *PaymentStatusReportV02) AddOriginalGroupInformationAndStatus() *model.OriginalGroupInformation1 {
	p.OriginalGroupInformationAndStatus = new(model.OriginalGroupInformation1)
	return p.OriginalGroupInformationAndStatus
}

func (p *PaymentStatusReportV02) AddTransactionInformationAndStatus() *model.PaymentTransactionInformation1 {
	newValue := new(model.PaymentTransactionInformation1)
	p.TransactionInformationAndStatus = append(p.TransactionInformationAndStatus, newValue)
	return newValue
}
