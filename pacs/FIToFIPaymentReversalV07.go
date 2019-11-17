package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700107 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.007.001.07 Document"`
	Message *FIToFIPaymentReversalV07 `xml:"FIToFIPmtRvsl"`
}

func (d *Document00700107) AddMessage() *FIToFIPaymentReversalV07 {
	d.Message = new(FIToFIPaymentReversalV07)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionPaymentReversal message is sent by an agent to the next party in the payment chain. It is used to reverse a payment previously executed.
// Usage
// The FIToFIPaymentReversal message is exchanged between agents to reverse a payment message that has been settled. The result will be a credit on the debtor account (when the reversed payment was a Direct Debit) or a debit on the creditor account (when the reversed payment was a Credit Transfer).
// The FIToFIPaymentReversal message may or may not be the follow-up of a payment message.
// The FIToFIPaymentReversal message refers to the original payment message by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentReversal message can be used in domestic and cross-border scenarios.
type FIToFIPaymentReversalV07 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader71 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *model.OriginalGroupHeader3 `xml:"OrgnlGrpInf,omitempty"`

	// Information concerning the original transactions, to which the reversal message refers.
	TransactionInformation []*model.PaymentTransaction81 `xml:"TxInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *FIToFIPaymentReversalV07) AddGroupHeader() *model.GroupHeader71 {
	f.GroupHeader = new(model.GroupHeader71)
	return f.GroupHeader
}

func (f *FIToFIPaymentReversalV07) AddOriginalGroupInformation() *model.OriginalGroupHeader3 {
	f.OriginalGroupInformation = new(model.OriginalGroupHeader3)
	return f.OriginalGroupInformation
}

func (f *FIToFIPaymentReversalV07) AddTransactionInformation() *model.PaymentTransaction81 {
	newValue := new(model.PaymentTransaction81)
	f.TransactionInformation = append(f.TransactionInformation, newValue)
	return newValue
}

func (f *FIToFIPaymentReversalV07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
