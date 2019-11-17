package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300106 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.06 Document"`
	Message *FIToFICustomerDirectDebitV06 `xml:"FIToFICstmrDrctDbt"`
}

func (d *Document00300106) AddMessage() *FIToFICustomerDirectDebitV06 {
	d.Message = new(FIToFICustomerDirectDebitV06)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionCustomerDirectDebit message is sent by the creditor agent to the debtor agent, directly or through other agents and/or a payment clearing and settlement system.
// It is used to collect funds from a debtor account for a creditor.
// Usage
// The FItoFICustomerDirectDebit message can contain one or more customer direct debit instructions.
// The FIToFICustomerDirectDebit message does not allow for grouping.
// The FItoFICustomerDirectDebit message may or may not contain mandate related information, i.e. extracts from a mandate, such as the MandateIdentification or DateOfSignature. The FIToFICustomerDirectDebit message must not be considered as a mandate.
// The FItoFICustomerDirectDebit message can be used in domestic and cross-border scenarios.
type FIToFICustomerDirectDebitV06 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader50 `xml:"GrpHdr"`

	// Set of elements providing information specific to the individual direct debit(s).
	DirectDebitTransactionInformation []*model.DirectDebitTransactionInformation20 `xml:"DrctDbtTxInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *FIToFICustomerDirectDebitV06) AddGroupHeader() *model.GroupHeader50 {
	f.GroupHeader = new(model.GroupHeader50)
	return f.GroupHeader
}

func (f *FIToFICustomerDirectDebitV06) AddDirectDebitTransactionInformation() *model.DirectDebitTransactionInformation20 {
	newValue := new(model.DirectDebitTransactionInformation20)
	f.DirectDebitTransactionInformation = append(f.DirectDebitTransactionInformation, newValue)
	return newValue
}

func (f *FIToFICustomerDirectDebitV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
