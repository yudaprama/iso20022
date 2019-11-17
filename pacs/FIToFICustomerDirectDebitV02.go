package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300102 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.02 Document"`
	Message *FIToFICustomerDirectDebitV02 `xml:"FIToFICstmrDrctDbt"`
}

func (d *Document00300102) AddMessage() *FIToFICustomerDirectDebitV02 {
	d.Message = new(FIToFICustomerDirectDebitV02)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionCustomerDirectDebit message is sent by the creditor agent to the debtor agent, directly or through other agents and/or a payment clearing and settlement system.
// It is used to collect funds from a debtor account for a creditor.
// Usage
// The FItoFICustomerDirectDebit message can contain one or more customer direct debit instructions.
// The FIToFICustomerDirectDebit message does not allow for grouping: the PaymentInformation block must be present once per occurrence of a DirectDebitTransactionInformation block.
// The FItoFICustomerDirectDebit message may or may not contain mandate related information, i.e. extracts from a mandate, such as the MandateIdentification or DateOfSignature. The FIToFICustomerDirectDebit message must not be considered as a mandate.
// The FItoFICustomerDirectDebit message can be used in domestic and cross-border scenarios.
type FIToFICustomerDirectDebitV02 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader34 `xml:"GrpHdr"`

	// Set of elements providing information specific to the individual direct debit(s).
	DirectDebitTransactionInformation []*model.DirectDebitTransactionInformation10 `xml:"DrctDbtTxInf"`
}

func (f *FIToFICustomerDirectDebitV02) AddGroupHeader() *model.GroupHeader34 {
	f.GroupHeader = new(model.GroupHeader34)
	return f.GroupHeader
}

func (f *FIToFICustomerDirectDebitV02) AddDirectDebitTransactionInformation() *model.DirectDebitTransactionInformation10 {
	newValue := new(model.DirectDebitTransactionInformation10)
	f.DirectDebitTransactionInformation = append(f.DirectDebitTransactionInformation, newValue)
	return newValue
}
