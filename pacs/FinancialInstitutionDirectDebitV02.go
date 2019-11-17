package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000102 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.02 Document"`
	Message *FinancialInstitutionDirectDebitV02 `xml:"FIDrctDbt"`
}

func (d *Document01000102) AddMessage() *FinancialInstitutionDirectDebitV02 {
	d.Message = new(FinancialInstitutionDirectDebitV02)
	return d.Message
}

// Scope:
// The FinancialInstitutionDirectDebit message is sent by an exchange or clearing house, or a financial institution, directly or through another agent, to the DebtorAgent.  It is used to instruct the DebtorAgent to move funds from one or more debtor(s) account(s) to one or more creditor(s), where both debtor and creditor are financial institutions.
//
// Usage:
// The FinancialInstitutionDirectDebit message is exchanged between agents and can contain one or more financial institution direct debit instruction(s) for one or more creditor(s). The FinancialInstitutionDirectDebit message can be used in domestic and cross-border scenarios.
//
type FinancialInstitutionDirectDebitV02 struct {

	// Common characteristics for all individual transactions included in the message.
	GroupHeader *model.GroupHeader63 `xml:"GrpHdr"`

	// Characteristics that apply to the credit side of the payment transaction(s) included in the message.
	CreditInstruction []*model.CreditTransferTransaction9 `xml:"CdtInstr"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *FinancialInstitutionDirectDebitV02) AddGroupHeader() *model.GroupHeader63 {
	f.GroupHeader = new(model.GroupHeader63)
	return f.GroupHeader
}

func (f *FinancialInstitutionDirectDebitV02) AddCreditInstruction() *model.CreditTransferTransaction9 {
	newValue := new(model.CreditTransferTransaction9)
	f.CreditInstruction = append(f.CreditInstruction, newValue)
	return newValue
}

func (f *FinancialInstitutionDirectDebitV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
