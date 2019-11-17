package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400101 struct {
	XMLName xml.Name          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.01 Document"`
	Message *PaymentReturnV01 `xml:"pacs.004.001.01"`
}

func (d *Document00400101) AddMessage() *PaymentReturnV01 {
	d.Message = new(PaymentReturnV01)
	return d.Message
}

// Scope
// The PaymentReturn message is sent by an agent to the previous agent in the payment chain to undo a payment previously settled.
// Usage
// The PaymentReturn message is exchanged between agents to return funds after settlement of credit transfer instructions (i.e. FIToFICustomerCreditTransfer message and FinancialInstitutionCreditTransfer message) or direct debit instructions (FIToFICustomerDirectDebit message).
// The PaymentReturn message should not be used between agents and non-financial institution customers. Non-financial institution customers will be informed about a debit or a credit on their account(s) through an Advice of Credit/Debit message and/or Statement message.
// The PaymentReturn message can be used to return single instructions or multiple instructions from one or different files.
// The PaymentReturn message can be used in domestic and cross-border scenarios.
// The PaymentReturn message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
type PaymentReturnV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader6 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *model.OriginalGroupInformation2 `xml:"OrgnlGrpInf,omitempty"`

	// Information concerning the original transactions, to which the return message refers.
	TransactionInformation []*model.PaymentTransactionInformation2 `xml:"TxInf,omitempty"`
}

func (p *PaymentReturnV01) AddGroupHeader() *model.GroupHeader6 {
	p.GroupHeader = new(model.GroupHeader6)
	return p.GroupHeader
}

func (p *PaymentReturnV01) AddOriginalGroupInformation() *model.OriginalGroupInformation2 {
	p.OriginalGroupInformation = new(model.OriginalGroupInformation2)
	return p.OriginalGroupInformation
}

func (p *PaymentReturnV01) AddTransactionInformation() *model.PaymentTransactionInformation2 {
	newValue := new(model.PaymentTransactionInformation2)
	p.TransactionInformation = append(p.TransactionInformation, newValue)
	return newValue
}
