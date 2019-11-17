package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02800101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.01 Document"`
	Message *FIToFIPaymentStatusRequestV01 `xml:"FIToFIPmtStsReq"`
}

func (d *Document02800101) AddMessage() *FIToFIPaymentStatusRequestV01 {
	d.Message = new(FIToFIPaymentStatusRequestV01)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionPaymentStatusRequest message is sent by the debtor agent to the creditor agent, directly or through other agents and/or a payment clearing and settlement system.  It is used to request a FIToFIPaymentStatusReport message containing information on the status of a previously sent instruction.
// Usage
// The FIToFIPaymentStatusRequest message is exchanged between agents to request status information about instructions previously sent. Its usage will always be governed by a bilateral agreement between the agents.
// The FIToFIPaymentStatusRequest message can be used to request information about the status (e.g. rejection, acceptance) of a credit transfer instruction, a direct debit instruction, as well as other intra-agent instructions (for example FIToFIPaymentCancellationRequest).
// The FIToFIPaymentStatusRequest message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentStatusRequest message can be used in domestic and cross-border scenarios.
//
//
type FIToFIPaymentStatusRequestV01 struct {

	// Set of characteristics shared by all individual transactions included in the status request message.
	GroupHeader *model.GroupHeader53 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status request message refers to.
	OriginalGroupInformation []*model.OriginalGroupInformation27 `xml:"OrgnlGrpInf,omitempty"`

	// Information concerning the original transaction, to which the status request message refers.
	TransactionInformation []*model.PaymentTransaction73 `xml:"TxInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *FIToFIPaymentStatusRequestV01) AddGroupHeader() *model.GroupHeader53 {
	f.GroupHeader = new(model.GroupHeader53)
	return f.GroupHeader
}

func (f *FIToFIPaymentStatusRequestV01) AddOriginalGroupInformation() *model.OriginalGroupInformation27 {
	newValue := new(model.OriginalGroupInformation27)
	f.OriginalGroupInformation = append(f.OriginalGroupInformation, newValue)
	return newValue
}

func (f *FIToFIPaymentStatusRequestV01) AddTransactionInformation() *model.PaymentTransaction73 {
	newValue := new(model.PaymentTransaction73)
	f.TransactionInformation = append(f.TransactionInformation, newValue)
	return newValue
}

func (f *FIToFIPaymentStatusRequestV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
