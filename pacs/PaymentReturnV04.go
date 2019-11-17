package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400104 struct {
	XMLName xml.Name          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.04 Document"`
	Message *PaymentReturnV04 `xml:"PmtRtr"`
}

func (d *Document00400104) AddMessage() *PaymentReturnV04 {
	d.Message = new(PaymentReturnV04)
	return d.Message
}

// Scope
// The PaymentReturn message is sent by an agent to the previous agent in the payment chain to undo a payment previously settled.
// Usage
// The PaymentReturn message is exchanged between agents to return funds after settlement of credit transfer instructions (i.e. FIToFICustomerCreditTransfer message and FinancialInstitutionCreditTransfer message) or direct debit instructions (FIToFICustomerDirectDebit message).
// The PaymentReturn message should not be used between agents and non-financial institution customers. Non-financial institution customers will be informed about a debit or a credit on their account(s) through a BankToCustomerDebitCreditNotification message ('notification') and/or BankToCustomerAccountReport/BankToCustomerStatement message ('statement').
// The PaymentReturn message can be used to return single instructions or multiple instructions from one or different files.
// The PaymentReturn message can be used in domestic and cross-border scenarios.
// The PaymentReturn message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
type PaymentReturnV04 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader54 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *model.OriginalGroupHeader2 `xml:"OrgnlGrpInf,omitempty"`

	// Information concerning the original transactions, to which the return message refers.
	TransactionInformation []*model.PaymentTransaction44 `xml:"TxInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentReturnV04) AddGroupHeader() *model.GroupHeader54 {
	p.GroupHeader = new(model.GroupHeader54)
	return p.GroupHeader
}

func (p *PaymentReturnV04) AddOriginalGroupInformation() *model.OriginalGroupHeader2 {
	p.OriginalGroupInformation = new(model.OriginalGroupHeader2)
	return p.OriginalGroupInformation
}

func (p *PaymentReturnV04) AddTransactionInformation() *model.PaymentTransaction44 {
	newValue := new(model.PaymentTransaction44)
	p.TransactionInformation = append(p.TransactionInformation, newValue)
	return newValue
}

func (p *PaymentReturnV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}
