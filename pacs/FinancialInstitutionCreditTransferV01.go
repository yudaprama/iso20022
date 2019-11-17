package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.01 Document"`
	Message *FinancialInstitutionCreditTransferV01 `xml:"pacs.009.001.01"`
}

func (d *Document00900101) AddMessage() *FinancialInstitutionCreditTransferV01 {
	d.Message = new(FinancialInstitutionCreditTransferV01)
	return d.Message
}

// Scope
// The FinancialInstitutionCreditTransfer message is sent by a debtor financial institution to a creditor financial institution, directly or through other agents and/or a payment clearing and settlement system.
// It is used to move funds from a debtor account to a creditor, where both debtor and creditor are financial institutions.
// Usage
// The FinancialInstitutionCreditTransfer message is exchanged between agents and can contain one or more credit transfer instructions where debtor and creditor are both financial institutions.
// The FinancialInstitutionCreditTransfer message does not allow for grouping: a CreditTransferTransactionInformation block must be present for each credit transfer transaction.
// The FinancialInstitutionCreditTransfer message can be used in domestic and cross-border scenarios.
//
type FinancialInstitutionCreditTransferV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader4 `xml:"GrpHdr"`

	// Set of elements providing information specific to the individual credit transfer(s).
	CreditTransferTransactionInformation []*model.CreditTransferTransactionInformation3 `xml:"CdtTrfTxInf"`
}

func (f *FinancialInstitutionCreditTransferV01) AddGroupHeader() *model.GroupHeader4 {
	f.GroupHeader = new(model.GroupHeader4)
	return f.GroupHeader
}

func (f *FinancialInstitutionCreditTransferV01) AddCreditTransferTransactionInformation() *model.CreditTransferTransactionInformation3 {
	newValue := new(model.CreditTransferTransactionInformation3)
	f.CreditTransferTransactionInformation = append(f.CreditTransferTransactionInformation, newValue)
	return newValue
}
