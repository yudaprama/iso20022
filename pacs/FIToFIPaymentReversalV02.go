package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.007.001.02 Document"`
	Message *FIToFIPaymentReversalV02 `xml:"FIToFIPmtRvsl"`
}

func (d *Document00700102) AddMessage() *FIToFIPaymentReversalV02 {
	d.Message = new(FIToFIPaymentReversalV02)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionPaymentReversal message is sent by an agent to the next party in the payment chain. It is used to reverse a payment previously executed.
// Usage
// The FIToFIPaymentReversal message is exchanged between agents to reverse a FIToFICustomerDirectDebit message that has been settled. The result will be a credit on the debtor account.
// The FIToFIPaymentReversal message may or may not be the follow-up of a CustomerDirectDebitInitiation message.
// The FIToFIPaymentReversal message refers to the original FIToFICustomerDirectDebit message by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentReversal message can be used in domestic and cross-border scenarios.
type FIToFIPaymentReversalV02 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader41 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *model.OriginalGroupInformation22 `xml:"OrgnlGrpInf"`

	// Information concerning the original transactions, to which the reversal message refers.
	TransactionInformation []*model.PaymentTransactionInformation29 `xml:"TxInf,omitempty"`
}

func (f *FIToFIPaymentReversalV02) AddGroupHeader() *model.GroupHeader41 {
	f.GroupHeader = new(model.GroupHeader41)
	return f.GroupHeader
}

func (f *FIToFIPaymentReversalV02) AddOriginalGroupInformation() *model.OriginalGroupInformation22 {
	f.OriginalGroupInformation = new(model.OriginalGroupInformation22)
	return f.OriginalGroupInformation
}

func (f *FIToFIPaymentReversalV02) AddTransactionInformation() *model.PaymentTransactionInformation29 {
	newValue := new(model.PaymentTransactionInformation29)
	f.TransactionInformation = append(f.TransactionInformation, newValue)
	return newValue
}
